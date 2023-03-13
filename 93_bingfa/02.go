package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	fmt.Println("main")
	time.Sleep(time.Second)
}

//执行了很多个匿名函数，因为一次循环的时间可以执行多个goroutine
//94
//9
//98
//99
//99
//99
//99
//100
//100