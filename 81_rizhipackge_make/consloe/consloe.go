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

// 其实Level就是levelStr，只不过类型不一样，l.Level就是主函数里传入的
// 给Logger定义一系列方法
func (l Logger) Debug(msg string) {
	if LogLevel(DEBUG) >= l.Level { //两个都是LogLevel类型的
		now := time.Now()
		TF := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] %s\n", TF, msg)
	}

}
func (l Logger) Info(msg string) {
	now := time.Now()
	TF := now.Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] %s\n", TF, msg)
}
func (l Logger) Warning(msg string) {
	now := time.Now()
	TF := now.Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] %s\n", TF, msg)
}
func (l Logger) Fatal(msg string) {
	now := time.Now()
	TF := now.Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] %s\n", TF, msg)
}
func (l Logger) Error(msg string) {
	now := time.Now()
	TF := now.Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] %s\n", TF, msg)
}
