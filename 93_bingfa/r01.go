package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func main() {
	wg1.Add(5)
	for i := 0; i < 5; i++ {
		//time.Sleep(time.Millisecond * 500)   1 2 3 4 5
		go func() {
			fmt.Println(i)
			wg1.Done()
		}()
		time.Sleep(time.Millisecond * 500) // 0 1 2 3 4
	}
	wg1.Wait()
	fmt.Println(666)
}
