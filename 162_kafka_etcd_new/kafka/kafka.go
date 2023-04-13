package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

//打开往kafka写日志的模块

type logDate struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer //声明一个全局的连接kafka的生产者client
	logDataChan chan *logDate
)

// Init 初始化client
func Init(addrs []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //发送完整数据类型leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //新选出一个partition
	config.Producer.Return.Successes = true                   //成功交付的消息将在success channel返回
	//连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed,err:", err)
		return
	}
	//初始化logDataChan
	logDataChan = make(chan *logDate, maxSize)
	//开启后台goroutine从通道中取数据发往kafka
	go SendToKafka()
	return
}

// 给外部暴露的一个函数，该函数只把日志数据发送到一个内部的channel中
func SendToChan(topic, data string) {
	msg := &logDate{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}

// 往kafka发送日志的函数
func SendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			//构造一个消息
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			//发送到kafka
			pid, offset, err := client.SendMessage(msg)
			fmt.Println("xxx")
			if err != nil {
				fmt.Println("send msg failed,err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
			fmt.Println("发送成功！")
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}
