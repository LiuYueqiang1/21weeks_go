package consloe

import (
	"fmt"
	"time"
)

// Logger 结构体
type Logger struct {
	Level LogLevel
}

// NewLog 建立构造函数调用这个结构体
func NewLog(levelStr string) Logger {
	level, err := pardeLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (l Logger) enable(loglevel LogLevel) bool {
	if loglevel >= l.Level { //l.Level==levelStr  true
		return true
	}
	return false
}

// 写一个记录日志的函数
func log(lv LogLevel, msg string) {
	now := time.Now()
	TF := now.Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] %s\n", TF, getLogString(lv), msg)
}

// 其实Level就是levelStr，只不过类型不一样，l.Level就是主函数里传入的
// 给Logger定义一系列方法
func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		//if LogLevel(DEBUG) >= l.Level { //两个都是LogLevel类型的
		//now := time.Now()
		//TF := now.Format("2006-01-02 15:04:05")
		//fmt.Printf("[%s] %s\n", TF, msg)
		log(DEBUG, msg)
	}
}
func (l Logger) Trace(msg string) {
	if l.enable(TRACE) {
		log(TRACE, msg)
	}
}
func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		log(INFO, msg)
	}

}
func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		log(WARNING, msg)
	}

}
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		log(FATAL, msg)
	}
}
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		log(ERROR, msg)
	}
}
