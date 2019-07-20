package main

import "fmt"

const (
	TCP = "tcp"
	UDP = "udp"
	WEBSOCKET = "ws"
)

type NetInfo struct {
	port string
	isTls  bool
}


var netInfos map[string]*NetInfo

func init() {
	netInfos = make(map[string]*NetInfo)
}

func HaveNet()bool  {
	if len(netInfos) > 0 {
		return true
	}
	return false
}


func CreateNet(potocol string,address string, tls_address string,
	cert_file string, key_file string)  {
	if potocol == "" || address == ""{
		fmt.Println("CreateNet parameter is invalid")
		return
	}

	isTls := false
	if tls_address != ""{
		isTls=true
	}
	netInfos[potocol] = &NetInfo{
		port:address,
		isTls:isTls,
	}

	if potocol == TCP {
		go StartSocketIO(address, tls_address,
			cert_file, key_file)
	}else if potocol == WEBSOCKET {
		go StartWSServer(address, tls_address,
			cert_file, key_file)
	}

}

