package consloe

import (
	"fmt"
	"path"
	"runtime"
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

// 定义一个比较数量级的方法
func (l Logger) enable(loglevel LogLevel) bool {
	if loglevel >= l.Level { //l.Level==levelStr  true
		return true
	}
	return false
}

// 查看执行的文件信息
func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("called is failed")
		return
	}
	fileName = runtime.FuncForPC(pc).Name()
	funcName = path.Base(file)
	return
}

// 写一个记录日志的函数
// 将记录日志的函数改成一个方法
func (l Logger) log(lv LogLevel, format string, a ...interface{}) {
	if l.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		funcName, fileName, lineNo := getInfo(3)
		now := time.Now()
		TF := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [%s] [文件名：%s 函数名：%s 行号：%d]%s\n", TF, getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

//在主函数中输入的不仅是一个字符串类型的，而是可以输入任意信息

// 其实Level就是levelStr，只不过类型不一样，l.Level就是主函数里传入的
// 给Logger定义一系列方法
func (l Logger) Debug(format string, a ...interface{}) {
	//if l.enable(DEBUG) {
	//	msg := fmt.Sprintf(format, a...)
	//	//if LogLevel(DEBUG) >= l.Level { //两个都是LogLevel类型的
	//	//now := time.Now()
	//	//TF := now.Format("2006-01-02 15:04:05")
	//	//fmt.Printf("[%s] %s\n", TF, msg)
	//	log(DEBUG, msg)
	//}
	l.log(DEBUG, format, a...)
}
func (l Logger) Trace(format string, a ...interface{}) {
	//if l.enable(TRACE) {
	msg := fmt.Sprintf(format, a...)
	l.log(TRACE, msg)
	//}
}
func (l Logger) Info(format string, a ...interface{}) {
	//if l.enable(INFO) {
	msg := fmt.Sprintf(format, a...)
	l.log(INFO, msg)
	//}

}
func (l Logger) Warning(format string, a ...interface{}) {
	//if l.enable(WARNING) {
	msg := fmt.Sprintf(format, a...)
	l.log(WARNING, msg)
	//}

}
func (l Logger) Fatal(format string, a ...interface{}) {
	//if l.enable(FATAL) {
	msg := fmt.Sprintf(format, a...)
	l.log(FATAL, msg)
	//}
}
func (l Logger) Error(format string, a ...interface{}) {
	//if l.enable(ERROR) {
	msg := fmt.Sprintf(format, a...)
	l.log(ERROR, msg)
	//}
}
