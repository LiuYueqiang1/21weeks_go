package main

import (
	"fmt"
	"sync"
)

// 先向通道存入数据,再从通道中获取数据
var a1 []int
var b1 chan int
var wg1 sync.WaitGroup

func main() {
	fmt.Println(b1) //<nil>
	b1 = make(chan int)
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		b1 <- 10
		fmt.Println("10发送到通道b中了", b1) //10发送到通道b中了 0xc00001c120
	}()

	x := <-b1
	fmt.Println("后台goroutine从通道b中获取到了", x) //后台goroutine从通道b中获取到了 10
	b1 = make(chan int, 16)
	fmt.Println(b1) //0xc00010a000
	wg1.Wait()
}
