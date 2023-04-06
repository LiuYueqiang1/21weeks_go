package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var cli = new(clientv3.Client)

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// 初始化ETCD的函数
func Init(addr string, timeout time.Duration) (err error) {
	// 创建etcd客户端
	cli, err = clientv3.New(clientv3.Config{
		//节点
		Endpoints: []string{addr},
		//5s钟都连不上就超时了
		DialTimeout: timeout,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	return
}

// 从ETCD中根据key获取配置项
func GetConf(key string) (logEntryConf []*LogEntry, err error) {
	// get 获取一个键的值
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	//Kvs多个键值对，一个个遍历出来
	for _, ev := range resp.Kvs {
		//fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed,err:%v\n", err)
			return
		}
	}
	return
}
