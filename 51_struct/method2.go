package main

import "fmt"

//给自定义类型加方法
//不能给别的暴力的类型添加方法，只能给自己包里的类型添加方法

type newInt int

// 方法    接受变量   方法名
func (n newInt) hello() {
	fmt.Println("这是一个int类型的方法")
}

// Question?    newInt100是什么
func main() {
	n1 := newInt(100)
	fmt.Println(n1)
	n1.hello()

	q2()
}

// 声明一个int32类型的变量x，它的值是10
func q1() {
	//方法1：
	var x int32
	x = 10
	//方法2：
	var x2 int32 = 10
	//方法3
	var x3 = int32(10)
	//方法4
	x4 := int32(10)
	fmt.Println(x, x2, x3, x4)
}
func q2() {
	//方法1
	var n1 newInt
	n1 = 100
	//方法2
	var n2 newInt = 100
	//方法3
	var n3 = newInt(100)
	//方法4
	n4 := newInt(100) //强制类型转换
	n4.hello()
	fmt.Println(n1, n2, n3, n4)
}
