package consloe

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// 在终端写相关的日志内容
type Loglevel uint16

const (
	UNKNOW Loglevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func pardeLoglevel(s string) (Loglevel, error) {
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
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("未知的错误类型")
		return UNKNOW, err
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}
