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
	sum int64
	job *job
}

var wg sync.WaitGroup

func f1(jobChan chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		jobChan <- newJob
		time.Sleep(time.Second)
	}
}
func f2(jobChan <-chan *job, resultChan chan<- *result) {
	defer wg.Done()

	for {
		n := <-jobChan
		nv := n.value
		//fmt.Println(nv)
		var sum = int64(0)
		for nv > 0 {
			sum = nv%10 + sum
			nv = nv / 10
		}
		x := &result{
			sum: sum,
			job: n,
		}
		resultChan <- x

	}
}
func main() {
	var jobChan = make(chan *job, 100)
	var resultChan = make(chan *result, 100)
	wg.Add(1)
	go f1(jobChan)
	wg.Add(24)
	for i := 1; i <= 24; i++ {
		go f2(jobChan, resultChan)
	}
	for result := range resultChan {
		fmt.Printf("随机数为 %d ,和为 %d\n", result.job.value, result.sum)
	}
	wg.Wait()
}
