package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 客户端
func main() {
	//1. 与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("dial \"tcp\",\"127.0.0.1:2000\" failed,err:", err)
		return
	}
	//2.发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请说话：")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}
