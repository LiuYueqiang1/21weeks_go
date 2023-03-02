package main

import "fmt"

//闭包
//闭包是一个函数，这个函数包含了他外部作用域的一个变量
//底层原理：
//1.函数可以作为返回值
//2.函数内部查找变量的顺序，先在自己内部找，找不到往外层找

// 要求让f1()调用f3()
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}
func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 传入 f2函数以及它的两个传入值 ，传出的是一个函数
func f3(f func(int, int), x, y int) func() {
	tep := func() {
		f(x, y) //设置一个匿名函数调用传入的函数并返回
	}
	return tep
}

func main() {
	ret := f3(f2, 100, 200) //把原来需要传递两个int类型的参数包装成一个不需要传参的函数
	f1(ret)
	//this is f1
	//this is f2
	//300
}

//func f3(x, y int) func() {
//	tmp := func() {
//		fmt.Println(x + y)
//	}
//	return tmp
//}
//
//func main() {
//	ret := f3(100, 200)
//	f1(ret)
//	//this is f1
//	//300
//	//f2未执行，执行的是f3的内置函数
//}
