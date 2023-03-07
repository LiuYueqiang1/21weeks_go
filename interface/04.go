package main

import "fmt"

// 接口的实现
type catt struct {
	name string
	feet int8
}

func (c catt) move() {
	fmt.Println("猫猫出击")
}
func (c catt) eat(food string) {
	fmt.Printf("%s爱吃吃%s...\n", c.name, food)
}

type animal interface {
	move()
	eat(fo string)
}

// ****
func hunr(a animal) {
	a.move()
	a.eat(string("猫粮"))
}
func main() {
	var aa animal
	fmt.Printf("接口的类型:%T\n", aa)
	aa = catt{
		name: "米粒",
		feet: 8,
	}
	aa.move()     //猫猫出击
	aa.eat("鱼罐头") //米粒爱吃吃鱼罐头...

	//***
	hunr(catt{
		name: "花花",
	}) //猫猫出击
	//花花爱吃吃猫粮...
}
