package main

import "21weeks/21weeks_go/82_test_rizhiku/mylogger"

// 自定义一个日志库
func main() {
	log := mylogger.Newlog()
	log.Debug("这是一条Debug日志")
	log.Info("这是一条Info日志")
	var l1 mylogger.Interfacer
	l1 = mylogger.Logger{}
	l1.Debug("这是一条Debug日志")
	l1.Info("这是一条Info日志")

}
