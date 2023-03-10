package main

import (
	"21weeks/21weeks_go/81_rizhipackge_make/consloe"
	"time"
)

func main() {
	log := consloe.NewLog()
	for {
		log.Debug("这是一条Debug日志")
		time.Sleep(2 * time.Second)
	}
}
