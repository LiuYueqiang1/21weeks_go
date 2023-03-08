package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入时如果有空格
func useScan() {
	var s string
	fmt.Println("请输入内容：")
	fmt.Scanln(&s) //读到空格或者enter就停止
	fmt.Printf("您输入的内容是%s\n", s)
}

// 使用bufio 可以获取整行
func useBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入内容")
	s, _ = reader.ReadString('\n')
	fmt.Printf("您输入的内容是：%s\n", s)
}
func main() {
	//useScan()
	//请输入内容：
	//a b c d e
	//您输入的内容是a
	useBufio()
	//a s d f s d
	//您输入的内容是：a s d f s d
}
