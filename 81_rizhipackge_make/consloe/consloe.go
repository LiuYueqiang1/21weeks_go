package consloe

import (
	"fmt"
	"time"
)

// 日志结构体
type Logger struct {
	Level Loglevel
}

// Logger的构造函数，调用Logger结构体
func NewLog(levelStr string) Logger {
	level, err := pardeLoglevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

// 构造一个比较的方法
func (l Logger) enable(logLevel Loglevel) bool {
	return logLevel >= l.Level
}

func log(lv Loglevel, msg string) {
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s:%s:%s],%s\n", now.Format("2006-01-02 15:04:05"), funcName, fileName, lineNo, msg)
}
func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		fmt.Printf("[%s],%s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		log(INFO, msg)
	}
}
func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		now := time.Now()
		fmt.Printf("[%s] ,%s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		fmt.Printf("[%s],%s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		now := time.Now()
		fmt.Printf("[%s],%s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
