package main

import "fmt"

// 闭包
// 每调用一次这个函数，+1
func intSeq() func() int {
	var i int = 0
	return func() int {
		i += 1
		return i
	}
}
func main() {
	nextInt := intSeq()           //闭包返回值是一个函数，用一个变量去接到它
	fmt.Printf("%T\n", intSeq())  //func() int
	fmt.Printf("%T\n", nextInt()) //int
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
