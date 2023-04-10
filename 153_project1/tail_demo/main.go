package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

// tail 读日志
func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 //重新打开（日志大小超出范围，重新打开）
		Follow:    true,                                 //是否跟随（继续读之前未读完的文件）
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
		MustExist: false,                                //文件不存在报错
		Poll:      true,                                 //轮询文件更改
	}
	tails, err := tail.TailFile(fileName, config) //用config配置项打开文件
	if err != nil {
		fmt.Println("tail file failed,err:", err)
		return
	}
	var (
		line *tail.Line //每一行
		ok   bool
	)
	for {
		line, ok = <-tails.Lines //通道里有值的话就读出来
		if !ok {
			fmt.Printf("tail file close reopen,filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}

		fmt.Println("line:", line.Text) //把读到的内容打印出来
	}
}
