package main

import "fmt"

type dog1 struct {
	name string
}

// 构造函数  调用结构体里面的东西 返回的是结构体名称
func newDog(name string) *dog1 {
	return &dog1{
		name: name,
	}
}

// 方法是作用于特定类型的函数
//
//	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//	   函数体
//	}
//
// 传入的   传出的
// 接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog1) wang() {
	fmt.Printf("%s:汪汪汪\n", d.name)
}
func main() {
	d1 := newDog("zzz") //返回的是结构体里面的东西，给d1
	d1.wang()
}
