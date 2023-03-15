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

// 服务端
func main() {
	//1.本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("strat tcp server on 127.0.0.1:2000 failed,err:", err)
		return
	}
	//2.等待客户端与我建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			return
		}
		go processConn(conn)
	}
}
