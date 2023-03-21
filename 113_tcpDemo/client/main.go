package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//建立与服务端的链接
//进行数据收发
//关闭链接

// 客户端
func main() {
	//1. 与server端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("dial tcp failed,err:", err)
		return
	}
	//2、发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入：")
		msg, _ := reader.ReadString('\n') //从终端读取输入
		msg = strings.TrimSpace(msg)      //TrimSpace返回字符串s的一个切片，删除所有前导和尾随的空白
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}
