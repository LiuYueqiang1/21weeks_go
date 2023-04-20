package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"review.com/logagent/conf"
	"review.com/logagent/kafka"
	"review.com/logagent/tail"
	"time"
)

var cfg = new(conf.SumConf)

func main() {
	//0.加载配置文件
	// 将配置文件加载出来映射到cfg对象里面
	err := ini.MapTo(cfg, "F:\\goland\\go_project\\21weeks\\21weeks_go\\157_re_logagent\\conf\\conf.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}
	//fmt.Println(cfg.KafkaConf.Address)
	//1、初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Println("init kafka failed,err:", err)
		return
	}
	fmt.Println("init kafka success!")
	//2、打开文件收集日志
	err = tail.Init(cfg.TailConf.Filename)
	if err != nil {
		fmt.Println("init taillog failed!,err:", err)
	}
	//3、开始业务，发送到kafka中
	fmt.Println("init taillog success")
	run()
}
func run() {
	for {
		select {
		case line := <-tail.ReadChan():
			kafka.SendTokafka(cfg.KafkaConf.Topic, line.Text)
		//2、发送到kafka
		default:
			time.Sleep(time.Second)
		}
	}
}
