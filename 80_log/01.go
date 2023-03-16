package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	filename, err := os.OpenFile("F:\\goland\\go_project\\21weeks\\21weeks_go\\80_log\\test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filename.Close()
	log.SetOutput(filename)
	for {
		log.Println("这是一条普通的日志")
		v := "狠狠普通"
		log.Printf("这是一条%s日志\n", v)
		//	log.Fatalln("这是一条会触发fatal的日志")
		//	log.Panicln("这是一条会触发panic的日志")
		time.Sleep(time.Second)
	}

}
