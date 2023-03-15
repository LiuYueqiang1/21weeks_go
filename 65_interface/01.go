package main

import "fmt"

type dog struct {
	name string
}
type cat struct {
	name string
}

// 方法
func (d dog) speak() {
	fmt.Printf("%v会汪汪汪~\n", d.name)
}
func (c cat) speak() {
	fmt.Println(("喵喵喵~"))
}

type speaker interface { //接口是一种类型
	speak() //接收到了什么方法
}

func da(x speaker) { //定义了一个名为da的函数，传入了一个变量，变量类型为接口类型
	x.speak() //这个接口类型的变量做了什么方法
}

func main() {
	var d1 dog
	d1.name = "大黄"
	var c1 cat
	//定义一个函数，传入一个接口，调用函数实现这个方法
	da(d1) //大黄会汪汪汪~
	da(c1) //喵喵喵~
	//将结构体传入接口，用接口实现方法
	var s1 speaker
	s1 = d1
	s1.speak()
	s1 = c1
	s1.speak()
	//直接用结构体实现方法
	d1.speak()
	c1.speak()
}
