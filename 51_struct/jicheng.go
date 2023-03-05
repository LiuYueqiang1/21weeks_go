package main

import "fmt"

// 继承
type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

type dog struct {
	feet   uint8
	animal //animal拥有的方法 和结构体，此时狗也拥有了
}

func (d dog) wang() {
	fmt.Printf("%s会汪汪汪\n", d.name)
}
func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			name: "岚牙",
		}, //类似于匿名嵌套结构体，但这是继承，可以使用animal的结构体
	}
	fmt.Println(d1) //{4 {岚牙}}
	d1.wang()       //岚牙会汪汪汪
	d1.move()       //继承自animal的方法    //岚牙会动
}
