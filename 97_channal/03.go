package main

import "fmt"

// 有缓冲区的通道
var b2 chan int

func main() {
	fmt.Println(b2)
	b2 = make(chan int, 2)
	b2 <- 10
	fmt.Println("10发送到通道中了")
	b2 <- 20
	fmt.Println("20发送到通道中了")
	x := <-b2
	fmt.Println("从通道b中接收到了", x)
}
