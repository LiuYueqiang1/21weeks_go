package mylogger

import "fmt"

//向终端写日志

// 定义Logger结构体
type Logger struct {
}

// 给Logger建立构造函数去调用Logger
func Newlog() Logger {
	return Logger{}
}

// 给Logger定义一系列方法
func (l Logger) Debug(msg string) {
	fmt.Println(msg)
}
func (l Logger) Info(msg string) {
	fmt.Println(msg)
}
func (l Logger) Warning(msg string) {
	fmt.Println(msg)
}
func (l Logger) Error(msg string) {
	fmt.Println(msg)
}
func (l Logger) Fatal(msg string) {
	fmt.Println(msg)
}

type Interfacer interface {
	Debug()
	Info()
}
