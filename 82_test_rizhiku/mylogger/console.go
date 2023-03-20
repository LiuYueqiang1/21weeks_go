package mylogger

import (
	"fmt"
	"time"
)

//向终端写日志

// 定义Logger结构体
type ConsoleLogger struct {
	Lever LogLevel
}

// 给Logger建立构造函数去调用Logger
func NewConsloelogger(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr) //输入一个String类型的，返回一个Loglevel类型的
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Lever: level,
	}
}

func (c ConsoleLogger) enable(loglevel LogLevel) bool {
	return loglevel >= c.Lever //c.Lever：写入的
}

// 写一个记日志的方法
func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...) //Sprintf根据格式说明符格式化并返回结果字符串
		now := time.Now()
		TF := now.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [文件名:%s 函数名:%s 行号:%d] %s\n", TF, getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

// 给Logger定义一系列方法
func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}
func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}
func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}
func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
