package taillog

import (
	"fmt"
	"kafka_etcd_new.com/kenew/etcd"
	"time"
)

var tskMgr *tailLogMgr

// 有很多taillog对象，管理taillog对象的，每一个对应着一个path，打开一个文件，对应一个tailObj
type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntryConf, //把当前的日志收集项保存起来
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), //无缓冲区的通道
	}
	for _, logEntry := range logEntryConf {
		//conf: *etcd.LogEntry
		NewTailTask(logEntry.Path, logEntry.Topic)

	}
	go tskMgr.run()
}

// 监听自己的newConfChan，有了新的配置过来之后就做对应的处理
// 配置新增
// 配置删除
// 配置变更
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			//1、配置新增
			//2、配置删除
			//3、配置变更
			fmt.Println("新的配置来了！", newConf)
		default:
			time.Sleep(time.Second)
		}

	}
}

// 向外暴露一个函数,向tskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
