package main

import (
	"fmt"
	"sync"
)

// 先从通道中获取数据，再向通道存入数据
var a []int
var b chan int
var wg sync.WaitGroup

func main() {
	fmt.Println(b) //<nil>
	b = make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中获取到了", x) //后台goroutine从通道b中获取到了 10
	}()
	b <- 10
	fmt.Println("10发送到通道b中了") //10发送到通道b中了
	b = make(chan int, 16)
	fmt.Println(b) //0xc00010a000
	wg.Wait()
}
