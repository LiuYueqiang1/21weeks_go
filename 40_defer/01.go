package main

import "fmt"

// Go语言中函数的return不是原子操作，在底层是分为两步来执行
//
// 第一步：返回值赋值
//
// 函数中如果存在defer，那么defer执行的时机实在第一步和第二步之间
//
// 第二步：真正的RET返回
func f1() int {
	x := 5
	defer func() {
		x++ //修改的是x，不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++ //函数传参 改变的是函数中的副本
	}(x)
	return 5 //返回值= x = 5
}
func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5
}

// 传入一个x的指针到匿名函数中
func f6() (x int) {
	defer func(x *int) {
		(*x)++ //跟f4相比，函数传入的是指针地址，修改的也是地址，故会发生变化
	}(&x)
	return 5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2()) //6
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5()) //5
	fmt.Println(f6()) //6
}
