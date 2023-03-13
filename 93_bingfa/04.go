package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		r1 := rand.Int() //int64
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}
func f1(i int) {
	defer wg.Done() //计数器-1
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) //计数器+1
		//如果计数器变为零，则释放在Wait上阻塞的所有goroutine
		go f1(i)
	}
	wg.Wait() //等待直到counter变为0
}
