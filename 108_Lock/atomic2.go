package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 原子操作版
type AtomicCounter struct {
	counter int64
}

func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}
func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter) //读值
}

type Counter2 interface {
	Inc()
	Load() int64
}

func test2(c Counter2) {
	var wg sync.WaitGroup
	strat := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(strat))
}
func main() {
	c3 := AtomicCounter{}
	test(&c3)
}
