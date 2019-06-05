package main

import (
	"bytes"
	"encoding/binary"
	engineio "github.com/googollee/go-engine.io"
	"github.com/gorilla/websocket"
	"io"
	"net"
	 "errors"
)
import "time"
import log "github.com/golang/glog"

type Client struct {
	conn   interface{}
	appid  int64
	connID uint64
	closed int32
	public_ip int32
	wt     chan *Message
}

type Message struct {
	cmd  int
	seq  int
	version int
	flag int

	body interface{}
}

func NewClient(conn interface{}) *Client {
	client := new(Client)

	//初始化Connection
	client.conn = conn // conn is net.Conn or engineio.Conn

	if net_conn, ok := conn.(net.Conn); ok {
		addr := net_conn.LocalAddr()
		if taddr, ok := addr.(*net.TCPAddr); ok {
			ip4 := taddr.IP.To4()
			client.public_ip = int32(ip4[0]) << 24 | int32(ip4[1]) << 16 | int32(ip4[2]) << 8 | int32(ip4[3])
		}
	}

	return client
}
const CLIENT_TIMEOUT = (60 * 6)
// 根据连接类型获取消息
func (client *Client) read() *Message {
	if conn, ok := client.conn.(net.Conn); ok {
		conn.SetReadDeadline(time.Now().Add(CLIENT_TIMEOUT * time.Second))
		return ReceiveClientMessage(conn)
	} else if conn, ok := client.conn.(engineio.Conn); ok {
		return ReadEngineIOMessage(conn)
	} else if conn, ok := client.conn.(*websocket.Conn); ok {
		conn.SetReadDeadline(time.Now().Add(CLIENT_TIMEOUT * time.Second))
		return ReadWebsocketMessage(conn)
	}
	return nil
}

//接受客户端消息(external messages)
func ReceiveClientMessage(conn io.Reader) *Message {
	return ReceiveLimitMessage(conn, 32*1024, true)
}

func ReceiveLimitMessage(conn io.Reader, limit_size int, external bool) *Message {
	buff := make([]byte, 12)
	_, err := io.ReadFull(conn, buff)
	if err != nil {
		log.Info("sock read error:", err)
		return nil
	}

	length, seq, cmd, version, flag := ReadHeader(buff)
	if length < 0 || length >= limit_size {
		log.Info("invalid len:", length)
		return nil
	}


	buff = make([]byte, length)
	_, err = io.ReadFull(conn, buff)
	if err != nil {
		log.Info("sock read error:", err)
		return nil
	}

	message := new(Message)
	message.cmd = cmd
	message.seq = seq
	message.version = version
	message.flag = flag
	message.body = buff
	return message
}

func WriteHeader(len int32, seq int32, cmd byte, version byte, flag byte, buffer io.Writer) {
	binary.Write(buffer, binary.BigEndian, len)
	binary.Write(buffer, binary.BigEndian, seq)
	t := []byte{cmd, byte(version), flag, 0}
	buffer.Write(t)
}

func ReadHeader(buff []byte) (int, int, int, int, int) {
	var length int32
	var seq int32
	buffer := bytes.NewBuffer(buff)
	binary.Read(buffer, binary.BigEndian, &length)
	binary.Read(buffer, binary.BigEndian, &seq)
	cmd, _ := buffer.ReadByte()
	version, _ := buffer.ReadByte()
	flag, _ := buffer.ReadByte()
	return int(length), int(seq), int(cmd), int(version), int(flag)
}

func WriteMessage(w *bytes.Buffer, msg *Message) {
	body := msg.body.([]byte)
	WriteHeader(int32(len(body)), int32(msg.seq), byte(msg.cmd), byte(msg.version), byte(msg.flag), w)
	w.Write(body)
}

func SendMessage(conn io.Writer, msg *Message) error {
	buffer := new(bytes.Buffer)
	WriteMessage(buffer, msg)
	buf := buffer.Bytes()
	n, err := conn.Write(buf)
	if err != nil {
		log.Info("sock write error:", err)
		return err
	}
	if n != len(buf) {
		log.Infof("write less:%d %d", n, len(buf))
		return errors.New("write less")
	}
	return nil
}

func ReadWebsocketMessage(conn *websocket.Conn) *Message {
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Info("read websocket err:", err)
		return nil
	}
	if messageType == websocket.BinaryMessage {
		return ReadBinaryMesage(p)
	} else {
		log.Error("invalid websocket message type:", messageType)
		return nil
	}
}
func ReadBinaryMesage(b []byte) *Message {
	reader := bytes.NewReader(b)
	return ReceiveClientMessage(reader)
}



func (client *Client) Read() {
	for {

		t1 := time.Now().Unix()
		msg := client.read()
		t2 := time.Now().Unix()
		if t2 - t1 > 6*60 {
			log.Infof("client:%d socket read timeout:%d %d", client.uid, t1, t2)
		}
		if msg == nil {
			client.HandleClientClosed()
			break
		}

		client.HandleMessage(msg)
		t3 := time.Now().Unix()
		if t3 - t2 > 2 {
			log.Infof("client:%d handle message is too slow:%d %d", client.connID, t2, t3)
		}
	}
}
func (client *Client) HandleMessage(msg *Message) {
	log.Info("msg cmd:", Command(msg.cmd))

}


func (client *Client) RemoveClient() {
	route := app_route.FindRoute(client.appid)
	if route == nil {
		log.Warning("can't find app route")
		return
	}
	route.RemoveClient(client)
}

func (client *Client) HandleClientClosed() {
	client.RemoveClient()
	//quit when write goroutine received
	client.wt <- nil
}

func (client *Client) AddClient() {
	route := app_route.FindOrAddRoute(client.appid)
	route.AddClient(client)
}



func (client *Client) Write() {
	running := true

	//发送在线消息
	for running {
		select {
		case msg := <-client.wt:
			if msg == nil {
				client.close()
				running = false
				log.Infof("client:%d socket closed", client.uid)
				break
			}

			seq++

			//以当前客户端所用版本号发送消息
			vmsg := &Message{msg.cmd, seq, client.version, msg.flag, msg.body}
			client.send(vmsg)
		}
	}

	//等待200ms,避免发送者阻塞
	t := time.After(200 * time.Millisecond)
	running = true
	for running {
		select {
		case <- t:
			running = false
		case <- client.wt:
			log.Warning("msg is dropped")
		}
	}

	log.Info("write goroutine exit")
}

func (client *Client) Run() {
	go client.Write()
	go client.Read()
}

