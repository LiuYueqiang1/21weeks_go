package main

import (
	"fmt"
	"sync"
)

var x = 0
var wg sync.WaitGroup
var lock sync.Mutex //互斥锁
func add() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		//修改之前加锁
		lock.Lock()
		x += 1
		//修改完成后解锁
		lock.Unlock()
	}
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
