package main

import (
	"177logtransfer.com/log/conf"
	"177logtransfer.com/log/es"
	"177logtransfer.com/log/kafka"

	"fmt"
	"gopkg.in/ini.v1"
)

//log transfer
//将日志数据从kafka取出来发往ES

func main() {
	//0、加载配置文件
	var cfg = new(conf.LogTransfer)                                                                      //在函数种修改变量需要传指针，配置文件对应的结构体中需要设置tag，特别时嵌套的结构体
	err := ini.MapTo(cfg, "F:\\goland\\go_project\\21weeks\\21weeks_go\\177_Logtransfer\\conf\\cfg.ini") //将一个变量传到函数里面，可以修改cfg变量的值，所以要new
	if err != nil {
		fmt.Printf("init config failed,err:%v\n", err)
		return
	}
	fmt.Println(cfg)
	//1、初始化kafka
	//初始化一个es连接的client
	//对外提供一个往ES写入数据的一个函数
	err = es.Init(cfg.EScfg.Address)
	if err != nil {
		fmt.Println("init es client failed,err:", err)
		return
	}
	fmt.Println("init es success")
	//连接kafka，创建分区消费者
	//每个分区消费者分别取出数据，通过sendToEs（）将数据发往es
	err = kafka.Init([]string{cfg.Kafka.Address}, cfg.Kafka.Topic)
	if err != nil {
		fmt.Println("init kafka failed,err:", err)
		return
	}
	//1、从kafka去日志数据

	//2、发往ES

}
