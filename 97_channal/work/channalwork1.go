package main

import (
	"fmt"
	"sync"
)

// 生成一百个数，存入通道1中
// 从通道1中取出，计算它的平方，存入通道2中
var wg sync.WaitGroup

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}
func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch2)
}
func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	wg.Add(2)
	go f1(ch1)
	go f2(ch1, ch2)
	for ret := range ch2 {
		fmt.Println(ret)
	}
	wg.Wait()
}
