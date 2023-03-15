package main

import "fmt"

type dog1 struct {
	name string
}
type cat1 struct {
	name string
}

// 方法
func (d dog1) speak() {
	fmt.Printf("%v会汪汪汪~\n", d.name)
}
func (c cat1) speak() {
	fmt.Println(("喵喵喵~"))
}

type speaker1 interface { //接口是一种类型
	speak() //接收到了什么方法
}

func da1(x speaker1) { //定义了一个名为da的函数，传入了一个变量，变量类型为接口类型
	x.speak() //这个接口类型的变量做了什么方法
}

func main() {
	var d1 dog1
	d1.name = "大黄"
	var c1 cat1
	da1(d1) //大黄会汪汪汪~
	da1(c1) //喵喵喵~

	var ss speaker1
	ss = d1
	//ss = c1
	fmt.Println(ss)
}
