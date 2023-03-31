package main

import (
	"fmt"
	"sync"
	"time"
)

// demo1 通道误用导致的bug
func demo1() {
	wg := sync.WaitGroup{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			for {
				task, ok := <-ch
				if ok == false {
					break
				}
				// 这里假设对接收的数据执行某些操作
				fmt.Println(task, j)
				time.Sleep(time.Millisecond * 500)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
func demo2() {
	wg := sync.WaitGroup{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	defer close(ch)

	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			for {
				select {
				case task := <-ch:
					fmt.Println(task)
				default:
					break
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
func main() {
	demo1()
}
