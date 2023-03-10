package main

import (
	test "69_package/calc"
	"fmt"
)

func init() {
	fmt.Println("main.init()")
}

func main() {
	fmt.Println("主函数")
	a := test.Name
	b, c := test.Add(5, 3)
	d := test.Sub(10, 6)
	fmt.Printf("calc包里面的Name：%v\n", a)
	fmt.Printf("调用calc包里面的Add函数：%v,%v\n", b, c)
	println(d)
}

//init函数按照插入顺序的逆序执行，因为它在包中自动首先执行，多用来做一些初始化的操作
