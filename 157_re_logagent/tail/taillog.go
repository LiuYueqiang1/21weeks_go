package tail

import (
	"fmt"
	"github.com/hpcloud/tail"
)

var (
	tailObj *tail.Tail
)

func Init() (err error) {
	config := tail.Config{
		ReOpen:    true,                                 //重新打开（日志大小超出范围，重新打开）
		Follow:    true,                                 //是否跟随（继续读之前未读完的文件）
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
		MustExist: false,                                //文件不存在报错
		Poll:      true,                                 //轮询文件更改
	}
	fileName := "F:\\goland\\go_project\\21weeks\\21weeks_go\\157_re_logagent\\157my.log"
	tailObj, err = tail.TailFile(fileName, config) //用config配置项打开文件
	if err != nil {
		fmt.Println("tail file failed,err:", err)
		return
	}
	return
}
