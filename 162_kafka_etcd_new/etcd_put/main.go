package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	//导入的包，v3表示版本
	// 创建etcd客户端
	cli, err := clientv3.New(clientv3.Config{
		//节点
		Endpoints: []string{"localhost:2379"},
		//5s钟都连不上就超时了
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put 创建一个键值
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"c/temp/.log","topic":"web_log"},{"path":"d/xxx/redix.log","topic":"web_log"}]`
	_, err = cli.Put(ctx, "/xxx", value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
}
