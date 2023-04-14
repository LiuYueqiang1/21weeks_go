package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var cli = new(clientv3.Client)

// 需要收集的日志的配置信息
type LogEntry struct {
	Path  string `json:"path"`  //日志存放的路径
	Topic string `json:"topic"` //日志要发往kafka中的哪个Topic
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
	//defer cli.Close()
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

// etcd watch
func WatchConf(key string, newConfCh chan<- []*LogEntry) {
	rch := cli.Watch(context.Background(), key) // <-chan WatchResponse
	//从通道中尝试取值（监视的信息）
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			//通知taillog.tskMgr
			//1、先判断操作的类型
			var newConf []*LogEntry
			if ev.Type != clientv3.EventTypeDelete {
				//如果是删除操作，手动传递一个空的配置项
		
				err := json.Unmarshal(ev.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("unmarshal failed,err:%v\n", err)
					continue
				}
			}
			fmt.Printf("get new conf:%v\n", newConf)
			newConfCh <- newConf

		}
	}
}
