package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// 基于sarama第三方库开发的kafka client  读取
// 将读取到的日志发送到kafka中
func main() {
	//配置信息
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完整数据类型leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition
	config.Producer.Return.Successes = true                   //成功交付的消息将在success channel返回
	//构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log" //Topic名称
	msg.Value = sarama.StringEncoder("this is a test log")
	//连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config) //使用给定的代理地址和配置config创建一个新的SyncProducer
	if err != nil {
		fmt.Println("producer closed,err:", err)
		return
	}
	fmt.Println("连接kafka成功")
	defer client.Close()
	//发送消息
	// 它将返回所生成消息的分区和偏移量，如果消息未能生成，则返回一个错误
	pid, offset, err := client.SendMessage(msg) //SendMessage生成一个给定的消息，只有当它成功或失败地生成时才返回。
	if err != nil {
		fmt.Println("send msg failed,err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset) //分区和索引位置
	fmt.Println("发送成功！")
}
