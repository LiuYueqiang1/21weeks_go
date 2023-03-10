package main

import (
	"21weeks/21weeks_go/81_rizhipackge_make/consloe"
	"fmt"
	"time"
)

func main() {
	log := consloe.NewLog("info")
	for i := 0; i <= 10; i++ {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Waining日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		fmt.Println("---------------------", i)
		time.Sleep(2 * time.Second)
	}
}
