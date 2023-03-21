package main

import (
	"fmt"
	"net"
)

func processConn(conn net.Conn) {
	//3.与客户端通信
	//	var tmp [128]byte
	var tymp = make([]byte, 128)
	for {
		n, err := conn.Read(tymp)
		if err != nil {
			fmt.Println("read from conn failed,err:", err)
			return
		}
		fmt.Println(string(tymp[:n]))
	}

}

//监听端口
//接收客户端请求建立链接
//创建goroutine处理链接。

// 服务端
// 一个服务器可以与多个用户端建立连接
func main() {
	//1.本地端口启动服务,监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:2000 failed,err:", err)
		return
	}
	//2、等待客户端与我建立连接，接受用户端信息
	for { //一直建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			return
		}
		//3、输出用户端信息
		//var tmp [128]byte
		//for { //循环读取，但是无法实现多个通信，没有办法一次读取多个客户端的信息，所以每新开一个用户就开一个goroutine，
		//	// 但是建立连接不需要开启多个goroutine，因为都是从一个端口监听的
		//	n, err := conn.Read(tmp[:])
		//	if err != nil {
		//		fmt.Println("read from conn failed,err:", err)
		//		return
		//	}
		//	fmt.Println(string(tmp[:n]))
		//}
		go processConn(conn) //每新开一个用户就开一个goroutine
	}
}
