package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
	LogChan chan string
)

func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
		MustExist: false,                                //文件不存在不报错
		Poll:      true,                                 //轮询文件更改
	}
	tailObj, err = tail.TailFile(fileName, config) //以配置项打开文件
	if err != nil {
		fmt.Println("tail file failed,err:", err)
		return
	}
	return
}

// 只读的单项通道（只要通道里有值就读出来）
func ReadChan() <-chan *tail.Line { //类型而已
	return tailObj.Lines
}
