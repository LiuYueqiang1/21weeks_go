package main

import (
	"fmt"
	"time"
)

// 给匿名函数传入循环的参数
// 这样的话一次只能调用一个
func main() {
	for i := 0; i < 100; i++ {
		go func(a int) {
			fmt.Println(a)
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)
}
