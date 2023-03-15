package main

import "fmt"

type baoshijie struct {
	brand string
}
type falali struct {
	brand string
}

func (b baoshijie) pao() {
	fmt.Printf("%s的速度是700迈\n", b.brand)
}
func (f falali) pao() {
	fmt.Printf("%s的速度是7000迈\n", f.brand)
}

type paoche interface { //定义接口
	pao()
}

func drive(p paoche) { //执行方法
	p.pao()
}
func main() {
	b1 := baoshijie{
		brand: "保时捷",
	}
	f1 := falali{
		brand: "法拉利",
	}
	b1.pao()
	f1.pao()
}
