package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后创建一个主goroutine去执行
func main() {
	for i := 0; i < 100; i++ {
		go hello(i) //开启一个单独的goroutine去执行hello函数
	}
	fmt.Println("main")
	time.Sleep(time.Second)
	//main函数结束了 由main函数启动的goroutine也结束了
}
