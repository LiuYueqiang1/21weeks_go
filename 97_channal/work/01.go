package main

//使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。
//开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
//开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
//主 goroutine 从resultChan取出结果并打印到终端输出

var jobChan = make(chan int64, 100)
var resultChan = make(chan int64, 100)

//函数f1（）循环生成int64类型的随机数，发送到jobChan中

//函数f2（）从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
