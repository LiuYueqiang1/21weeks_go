package main

import "fmt"

// 一个类型实现多个接口
type cat2 struct { //一个类型
	name string
}

func (c *cat2) move() {
	fmt.Println("猫猫出击")
}
func (c *cat2) speak() {
	fmt.Println("喵喵叫")
}

type movee interface { //一个接口对应一个方法
	move()
}

type speakk interface { //一个接口对应一个方法，一个接口也可以对应多个方法
	speak()
}

type sum interface { //使用一个接口  调用两个接口的类型名字
	speakk
	movee
}

func main() {
	var m1 movee
	m1 = &cat2{
		"米粒",
	}
	m1.move()
	var s1 speakk
	s1 = &cat2{
		name: "米粒",
	}
	s1.speak()

	//*****
	var ss sum //调用两个接口 放入一个接口中
	ss = &cat2{
		name: "米粒",
	}
	ss.move()
	ss.speak()
}
