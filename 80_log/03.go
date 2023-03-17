package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("F:\\goland\\go_project\\21weeks\\21weeks_go\\80_log\\test2.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	//输出微秒级别的时间，文件名+行号+日期
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("111这是一条很普通的日志。")
	log.SetPrefix("[小王子]")
	log.Println("这是一条很普通的日志。")

	//log 标准库 中Logger的构造函数
	//func New(out io.Writer, prefix string, flag int) *Logger
	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。") //打印到终端
	log.Println("这是自定义的logger记录的日志。")    //打印到自己打开的文件里
}
