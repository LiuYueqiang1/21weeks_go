package main

import (
	"fmt"
	"strings"
)

// 切割字符串测试
func Split(str, sep string) (result []string) {
	i := strings.Index(str, sep)
	for i >= 0 {
		result = append(result, str[:i])
		str = str[i+len(sep):]
		i = strings.Index(str, sep)
	}
	result = append(result, str)
	return
}

func main() {
	ret := Split("abcdefg", "de")
	fmt.Println(ret)
}
