package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/*
所谓的大小端就是指字节序在内存中是如何存储的。大端指的就是把字节序的尾端（0xcd）放在高内存地址，
而小端指的就是把字节序的尾端（0xcd）放在低内存地址，所以正确的叫法应该是高尾端和低尾端。
【注】不管是大端法还是小端法存储，计算机在内存中存放数据的顺序都是从低地址到高地址，所不同的是首先 取低字节的数据存放在低地址 还是 取高字节数据存放在低地址。

著名的CPU两大派系，PowerPC系列采用大端（big endian）的方式存储数据，而X86系列则采用小端（little endian）方式存储数据。
很显然如果你的程序只运行在PowerPC系列的CPU上，你完全可以不管什么是little endian，但是如果你PowerPC上的程序要和X86上的程序打交道，
那么你就必须进行转换才能相互识别，不然那是要打架的。看见没有，计算机世界尽是与现实生活如此相似。


1、发送的时候使用：htons（l）
2、接受的时候使用：ntohs（l）


大部分小型机采用big endian，运行的是unix系统，也有一些小型机是little endian 系统，
如康柏的vms！intel的处理器是little endian 的，所以windows是little endian系统。
是什么字续，主要取决与处理器的处理顺序！


一般来说，除了intel 80x86系列处理器是小尾架构，绝大部分处理器均为大尾架构，如sparc系列/power系列/moto的68系列等。  网络字节顺序也是大端的。
在编解码时，尤其需要注意大小尾问题。在每处使用超过一个byte的地方，最好使用转换函数（hton*和ntoh*系列或自写均可）

所以，当你的通信软件要和其他机器上的通信软件（模块）通信时，凡是编解码等地方使用了超过1个字节的数据类型，都最好使用转换函数。
在部分socket处理中，也需要加上转换函数（如ipaddr结构的填写等处），另外一部分本身已经包含相关处理，就可以不用加。
*/

func main() {
	var i uint32 = 256 + 16 + 1

	// 小端
	b := make([]byte, 4)
	/*
		b[0] = byte(v)
		b[1] = byte(v >> 8)
		b[2] = byte(v >> 16)
		b[3] = byte(v >> 24)
	*/
	binary.LittleEndian.PutUint32(b, i)
	fmt.Printf("LittleEndian(%d)  b[0] b[1] b[2] b[3]:", i)
	for _, bin := range b {
		fmt.Printf("%02X ", bin)
	}
	fmt.Printf("\n")

	//大端
	fmt.Printf("BigEndian(%d)     b[3] b[2] b[1] b[0]:", i)
	binary.BigEndian.PutUint32(b, i)
	for _, bin := range b {
		fmt.Printf("%02X ", bin)
	}
	fmt.Printf("\n")

	//[]byte 2 uint32
	bytesBuffer := bytes.NewBuffer(b)
	var j uint32
	binary.Read(bytesBuffer, binary.BigEndian, &j)
	fmt.Println("j = ", j)

}
