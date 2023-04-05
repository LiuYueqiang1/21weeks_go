package main

import (
	"example.com/logagent/kafka"
	"example.com/logagent/taillog"
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"time"
)

var (
	err error
	cfg *ini.File
)

func run() {
	//1、读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			//2、发送到kafka
			kafka.SendToKafka(cfg.Section("kafka").Key("topic").String(), line.Text)
		default:
			time.Sleep(time.Second)
		}

	}
}

// logAgent入口程序
func main() {
	//0.加载配置文件
	cfg, err = ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	//init 不能写注释
	fmt.Println(cfg.Section("kafka").Key("address").String())
	fmt.Println(cfg.Section("kafka").Key("topic").String())
	fmt.Println(cfg.Section("taillog").Key("path").String())
	//1、初始化kafka连接
	err = kafka.Init([]string{cfg.Section("kafka").Key("address").String()})
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")
	//2、打开日志文件准备收集日志
	err = taillog.Init(cfg.Section("taillog").Key("path").String())
	if err != nil {
		fmt.Printf("init taillog failed,err:%v\n", err)
		return
	}
	fmt.Println("init taillog success")
	run()
}
