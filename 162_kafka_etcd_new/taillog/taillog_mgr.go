package taillog

import "kafka_etcd_new.com/kenew/etcd"

var tskMgr *taillogMgr

// 有很多taillog对象，管理taillog对象的，每一个对应着一个path，打开一个文件，对应一个tailObj
type taillogMgr struct {
	logEntry []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &taillogMgr{
		logEntry: logEntryConf, //把当前的日志收集项保存起来
	}
	for _, logEntry := range logEntryConf {
		//conf: *etcd.LogEntry
		NewTailTask(logEntry.Path, logEntry.Topic)

	}
}
