package taillog

import (
	"fmt"
	"kafka_etcd_new.com/kenew/etcd"
	"time"
)

// 165
var tskMgr *tailLogMgr

// 有很多taillog对象，管理taillog对象的，每一个对应着一个path，打开一个文件，对应一个tailObj
type tailLogMgr struct {
	logEntry    []*etcd.LogEntry     //是否修改
	tskMap      map[string]*TailTask //有多少个tailtask记下来
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
		//logEntry.Path
		//初始化的时候起了多少个tailtask 都要记下来，为了后续判断方便
		tailtask := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tskMap[mk] = tailtask

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
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					//原来就有，不需要操作
					continue
				} else {
					tailtask := NewTailTask(conf.Path, conf.Topic)
					t.tskMap[mk] = tailtask
				}
			}
			//找出原来t，logEntrt有，但是newConf中没有的，要删掉
			for _, c1 := range t.logEntry { // 从原来的
				isDelete := true
				for _, c2 := range newConf { //去新的配置中逐一进行比较
					if c2.Path == c1.Path && c2.Topic == c1.Topic { //原来有现在也有，什么都不做
						isDelete = false
						continue
					}
				}
				if isDelete {
					//把c1对应的这个tailObj给停掉
					mk := fmt.Sprintf("%s_%s", c1.Path, c1.Topic)
					t.tskMap[mk].cancelFunc()
				}
			}
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
