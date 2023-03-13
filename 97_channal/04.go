package main

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

func f1(ch1 chan int) {
	defer wg2.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}
func f2(ch1, ch2 chan int) {
	defer wg2.Done()
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
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg2.Add(2)
	go f1(a)
	go f2(a, b)
	wg2.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}
