package main

import (
	"21weeks/21weeks_go/82_test_rizhiku/mylogger"
	"time"
)

// 自定义一个日志库
func main() {
	log := mylogger.Newlog("Debug")
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Erroe日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second)
	}
}
