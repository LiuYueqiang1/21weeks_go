package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var client sarama.SyncProducer

func Init(addr []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	// 连接kafka
	client, err = sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	//defer client.Close()
	return
}
func SendTokafka(topic, data string) {
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	fmt.Println("xxx")
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	fmt.Println("发送到kafka成功！")
}
