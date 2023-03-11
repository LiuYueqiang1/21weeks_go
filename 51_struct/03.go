package main

import "fmt"

// 匿名字段
// 字段比较少比较简单
// 不常用
type person4 struct {
	string
	int
}
type person5 struct {
	name string
	age  int
}

type address struct {
	city string
	mail string
}

// 嵌套结构体
type company struct {
	name    string
	pe      person5
	address //匿名嵌套结构体
}

func main() {
	//	person5.name=  非法的吧

	p1 := person4{
		"米栗木",
		18,
	}
	fmt.Println(p1)
	c1 := company{
		name: "华强集团",
		pe: person5{
			"刘华强",
			28,
		},
	}
	fmt.Println(c1) //{华强集团 {刘华强 28}}
	c2 := company{
		name: "撒日朗",
		pe: person5{
			name: "华强",
			age:  28,
		},
		address: address{
			"北京不知名水果摊",
			"保熟吗.com",
		},
	}
	fmt.Println(c2.pe.name) //普通嵌套结构体    //华强
	fmt.Println(c2.city)    //先在自己的结构体里查找该字段，找不到就去匿名嵌套的结构体中查找  //北京不知名水果摊
	fmt.Println(c2)         //{撒日朗 {华强 28} {北京不知名水果摊 保熟吗.com}}
}
