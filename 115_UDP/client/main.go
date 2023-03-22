package main

import (
	"fmt"
	"net"
)

func main() {
	connUDP, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务器失败,err:", err)
		return
	}
	defer connUDP.Close()
	sentDate := []byte("hello server")
	_, err = connUDP.Write(sentDate) //发送数据
	if err != nil {
		fmt.Println("client write failed,err:", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteUDP, err := connUDP.ReadFromUDP((data)) //接收数据
	if err != nil {
		fmt.Println("接收数据失败,err:", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteUDP, n)
}
