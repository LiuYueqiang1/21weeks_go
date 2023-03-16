package main

import (
	"bufio"
	"fmt"
	"os"
)

// 使用bufio读取整行输入
func main() {
	fmt.Println("请输入：")
	s := bufio.NewReader(os.Stdin)
	ss, err := s.ReadString('\n')
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("输入的结果为：")
	fmt.Println(ss)
}
