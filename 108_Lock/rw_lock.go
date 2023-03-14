package main

import (
	"fmt"
	"sync"
	"time"
)

var wg1 sync.WaitGroup
var lo sync.Mutex
var rw sync.RWMutex
var y = 0

func read() {
	defer wg1.Done()
	//lo.Lock()
	rw.RLock()
	fmt.Println(y)
	time.Sleep(time.Millisecond)
	//lo.Unlock()
	rw.RUnlock()
}
func write() {
	defer wg1.Done()
	//	lo.Lock()
	rw.RLock()
	y = y + 1
	time.Sleep(time.Millisecond * 5)
	//lo.Unlock()
	rw.RUnlock()
}
func main() {
	now := time.Now()
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go write()
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		wg1.Add(1)
		go read()
	}
	wg1.Wait()
	tsub := time.Now().Sub(now)
	fmt.Println(tsub)
}
