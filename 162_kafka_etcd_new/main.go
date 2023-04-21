package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"kafka_etcd_new.com/kenew/conf"
	"kafka_etcd_new.com/kenew/etcd"
	"kafka_etcd_new.com/kenew/kafka"
	"kafka_etcd_new.com/kenew/taillog"
	"kafka_etcd_new.com/kenew/utils"
	"sync"
	"time"
)

var (
	cfg = new(conf.Sumconfig)
)

// logAgent入口程序
func main() {
	// 将配置文件加载出来映射到cfg对象里面
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}

	//1、初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
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

	//初始化etcd后不用tail包从文件中读了，而是从etcd里拉取
	// 为了实现每个logagent都拉去自己独有的配置，所以要自己的IP地址作为区分
	ipStr, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)
	//2、1从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(etcdConfKey)

	if err != nil {
		fmt.Printf("etcd.GetConf failed,err:%v\n", err)
		return
	}
	fmt.Printf("get conf from etcd success,%v\n", logEntryConf)
	//2、2拍一个哨兵去监视日志收集项的变化（有变化及时通知我的logAgent实现加载配置）
	for index, value := range logEntryConf {
		fmt.Printf("index:%v  value:%v\n", index, value)
	}
	//读文件 发往kafka的topic
	//[{"path":"c/temp/.log","topic":"web_log"},{"path":"d/xxx/redix.log","topic":"web_log"}]
	//3、从path收集日志发往kafka的topic
	//（每一个path对应一个tailObj）
	//3、1循环每一个日志收集项，创建TailObj
	taillog.Init(logEntryConf) //因为NewConfChan访问了tskMgr的newConfChan，
	//这个channnel是在taillog.Init(logEntryConf)执行的初始化
	newConfChan := taillog.NewConfChan() //从taillog包中获取对外暴露的通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan) //哨兵发现最新的配置信息会通知上面那个通道
	wg.Wait()
	//3、2发往kafka
}
