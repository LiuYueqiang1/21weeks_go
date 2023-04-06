package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"kafka_etcd_new.com/kenew/conf"
	"kafka_etcd_new.com/kenew/etcd"
	"kafka_etcd_new.com/kenew/kafka"
	"time"
)

var (
	cfg = new(conf.Sumconfig)
)

//func run() {
//	//1、读取日志
//	for {
//		select {
//		case line := <-taillog.ReadChan():
//			//2、发送到kafka
//			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
//		default:
//			time.Sleep(time.Second)
//		}
//	}
//}

// logAgent入口程序
func main() {
	// 将配置文件加载出来映射到cfg对象里面
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}

	//1、初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init Kafka failed,err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")
	//2、初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed,err:%v\n", err)
		return
	}
	fmt.Println("init etcd success")
	//2、打开日志文件准备收集日志
	//err = taillog.Init(cfg.TaillogConf.Filename)
	//if err != nil {
	//	fmt.Printf("init taillog failed,err:%v\n", err)
	//	return
	//}
	//fmt.Println("init taillog success")
	//3、具体的业务
	//run()
}
