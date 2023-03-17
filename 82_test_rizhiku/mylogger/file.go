package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写代码
type FileLogger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

// FileLogger 的构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = f1.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return f1
}
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file dailed,err:%v\n", err)
		return err
	}
	errfileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file dailed,err:%v\n", err)
		return err
	}
	//日志文件已打开
	f.fileObj = fileObj
	f.errFileObj = errfileObj
	return nil
}

// 写一个记日志的方法
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...) //Sprintf根据格式说明符格式化并返回结果字符串
		now := time.Now()
		TF := now.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getInfo(3)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [文件名:%s 函数名:%s 行号:%d] %s\n", TF, getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			//如果要记录的日志大于等于ERROR级别，还需要再err日志中再记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [文件名:%s 函数名:%s 行号:%d] %s\n", TF, getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

// 写一个记日志的函数
func (f *FileLogger) enable(loglevel LogLevel) bool {
	return loglevel >= f.Level //f.Lever：写入的
}

// 给Logger定义一系列方法
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
