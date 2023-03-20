package main

import (
	"21weeks/21weeks_go/81_rizhipackge_make/consloe"
	"time"
)

func main() {
	log := consloe.NewLog("debug")
	//	fmt.Println(log)  1
	for {
		id := 10010
		name := "米栗木"
		log.Debug("这是一条Debug日志,id:%d,name:%s", id, name)
		log.Trace("这是一条Trace日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second)
	}

}
