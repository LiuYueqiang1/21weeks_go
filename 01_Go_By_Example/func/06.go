package main

import "fmt"

//方法

type rect struct {
	width, height int
}

//	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//		函数体
//	}
func (r *rect) area() int {
	return r.width * r.height
}
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
func main() {
	r := rect{
		width:  10,
		height: 5,
	}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	fmt.Println((&r).area())
	fmt.Println((&r).perim())
}
