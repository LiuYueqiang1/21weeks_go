package consloe

import (
	"errors"
	"strings"
)

// 给log分级别
type LogLevel uint16

const (
	UNKONW LogLevel = iota
	DEBUG           //1
	TRACE
	INFO
	WARNING
	ERROR
	FATAL //6
)

// 将主函数中输入的String类型转换为LogLevel类型
func pardeLogLevel(s string) (LogLevel, error) {
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
		return UNKONW, err
	}

}
