package main

import (
	"example.com/logagent/kafka"
	"example.com/logagent/taillog"
	"fmt"
	"time"
)

func run() {
	//1、读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			//2、发送到kafka
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}

	}
}

// logAgent入口程序
func main() {
	//1、初始化kafka连接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")
	//2、打开日志文件准备收集日志
	err = taillog.Init("F:\\goland\\go_project\\21weeks\\my.log")
	if err != nil {
		fmt.Printf("init taillog failed,err:%v\n", err)
		return
	}
	fmt.Println("init taillog success")
	run()
}
