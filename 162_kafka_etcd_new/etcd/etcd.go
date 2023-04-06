package etcd

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var cli = new(clientv3.Client)

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
