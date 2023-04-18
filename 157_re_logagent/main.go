package main

import (
	"fmt"
	"review.com/logagent/kafka"
	"review.com/logagent/tail"
)

func main() {
	//1、初始化kafka连接
	kafka.Init()
	fmt.Println("init kafka success!")
	//2、打开文件收集日志
	err := tail.Init()
	if err != nil {
		fmt.Println("init taillog failed!,err:", err)
	}
	//3、开始业务，发送到kafka中
	run()
}
func run() {
	kafka.SendTokafka()
}
