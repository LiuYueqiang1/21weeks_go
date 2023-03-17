package mylogger

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//向终端写日志

// 给log分级别
type LogLevel uint16 //这里必须是type，而不是var

const (
	UNKNON LogLevel = iota
	TRACE
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

// 定义Logger结构体
type Logger struct {
	Lever LogLevel
}

// 将构造函数里面string类型的转换为Logger类型
func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNON, err

	}
}

// 给Logger建立构造函数去调用Logger
func Newlog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr) //输入一个String类型的，返回一个Loglevel类型的
	if err != nil {
		panic(err)
	}
	return Logger{
		Lever: level,
	}
}

func (l Logger) enable(loglevel LogLevel) bool {
	return loglevel >= l.Lever //l.Lever：写入的
}

// 给Logger定义一系列方法
func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		TF := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [DEBUG] %s\n", TF, msg)
	}

}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		TF := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [INFO] %s\n", TF, msg)
	}
}
func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		now := time.Now()
		TF := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [WARNING] %s\n", TF, msg)
	}
}
func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		TF := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [ERROR] %s\n", TF, msg)
	}

}
func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		now := time.Now()
		TF := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [FATAL] %s\n", TF, msg)
	}
}
