package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。
//开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
//开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
//主 goroutine 从resultChan取出结果并打印到终端输出

type job struct {
	value int64
}
type result struct {
	job *job
	sum int64
}

// 初始化通道
var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

// 生成int64的随机数
func f1(ch1 chan<- *job) {
	defer wg.Done()
	for {
		a := rand.Int63()
		newjob := &job{
			value: a,
		}
		ch1 <- newjob
		time.Sleep(time.Millisecond * 500)
	}
}

// 取出和
// 如果要用通道作为返回值，直接给ch3，并return
// func f2(ch1 <-chan *job, ch2 chan<- *result) (ch3 chan <- *result) {
func f2(ch1 <-chan *job, ch2 chan<- *result) (ch3 chan<- *result) {
	defer wg.Done()
	for {
		sum := int64(0)
		a1 := <-ch1
		n := a1.value
		//不能直接传入a1.value，此时传入的是指针，直接没有了
		//先赋值给n，然后用值去解决
		for n > 0 {
			sum = sum + n%10
			n = n / 10
		}
		//fmt.Println(a1.value)
		ret := &result{
			job: a1,
			sum: sum,
		}
		//ch3 <- ret
		//return ch3
		ch2 <- ret

	}

}

func main() {
	wg.Add(1)
	go f1(jobChan)
	for i := 0; i < 24; i++ {
		wg.Add(1)
		go f2(jobChan, resultChan)
	}
	for v := range resultChan {
		fmt.Println(v.job.value, v.sum)
	}
}
