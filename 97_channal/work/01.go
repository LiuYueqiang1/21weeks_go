package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。
// 开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
// 开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
// 主 goroutine 从resultChan取出结果并打印到终端输出
type job struct {
	value int64
}
type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wwg sync.WaitGroup

// 函数f1（）循环生成int64类型的随机数，发送到jobChan中
func ff1(ch1 chan<- *job) {
	defer wwg.Done()
	for {
		a1 := rand.Int63()
		//fmt.Println("接收success")
		newjob := &job{
			value: a1,
		}
		ch1 <- newjob
		time.Sleep(time.Millisecond * 500)
		//fmt.Println("接收success")
	}

}

// 函数f2（）从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
func ff2(ch1 <-chan *job, ch2 chan<- *result) {
	defer wwg.Done()
	for {
		a1 := <-ch1
		//if !ok {
		//	break
		//}
		//fmt.Println(a1)
		n := a1.value
		sum := int64(0)

		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: a1,
			sum: sum,
		}
		//fmt.Println(sum)
		ch2 <- newResult
	}

}
func main() {
	wwg.Add(1)
	go ff1(jobChan)
	//for result1 := range jobChan {
	//	fmt.Println(result1)
	//}
	wwg.Add(24)
	for i := 0; i < 24; i++ {
		go ff2(jobChan, resultChan)
	}

	for result := range resultChan {
		fmt.Println(result.job.value, result.sum)
	}
	wwg.Wait()
}
