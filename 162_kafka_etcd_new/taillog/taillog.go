package taillog

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"kafka_etcd_new.com/kenew/kafka"
)

// TailTask 一个日志收集的任务
type TailTask struct {
	path     string //收集哪个路径的日志
	topic    string //收集的日志保存到kafka下的哪个topic
	instance *tail.Tail
	//为了能实现退出t.run()
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	tailObj.init()
	return
}
func (t *TailTask) init() {
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
		MustExist: false,                                //文件不存在不报错
		Poll:      true,                                 //轮询文件更改
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed,err:", err)
		return
	}
	//当goroutine执行的函数推出的时候，goroutine就结束了
	go t.run() //直接去采集日志发送到kafka
}
func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task:%s_%s 结束了...\n", t.path, t.topic)
			return
		case line := <-t.instance.Lines:
			//kafka.SendToKafka(t.topic, line.Text) //函数调用函数（修改：放到通道里即可）
			//先把日志数据发送到一个通道中
			kafka.SendToChan(t.topic, line.Text)
			//kafka那个包中有单独的goroutine去收集日志数据发送到kafka
		}
	}
}

// 只读的单项通道
func (t *TailTask) ReadChan() <-chan *tail.Line {
	return t.instance.Lines
}
