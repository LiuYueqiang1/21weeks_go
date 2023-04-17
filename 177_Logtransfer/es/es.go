package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var client *elastic.Client
var ch chan *LogData

//初始化ES 准备接收kafka发来的数据

func Init(address string) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "heep://" + address
	}

	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		// Handle error
		return err
	}
	fmt.Println("connect to es success")
	ch = make(chan *LogData, 10000)
	go SendToES()
	return
}

func SendToESChan(msg *LogData) {
	ch <- msg
}

// sendToES 发送数据到ES
func SendToES() {

	//链式操作
	for {
		select {
		case msg := <-ch:
			put1, err := client.Index().
				Index(msg.Topic).
				BodyJson(msg).
				Do(context.Background())
			if err != nil {
				// Handle error
				panic(err)
			}
			fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}

}
