package main

import (
	"fmt"
	"reflect"
)

// 类型（Type）和种类（Kind）
type myInt int64

func reflecType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}
func main() {
	var a *float32
	var b myInt
	var c rune    //Rune是int32的别名，在所有方面都等价于int32。按照惯例，它被用来区分字符值和整数值。
	reflecType(a) //type: kind:ptr
	reflecType(b) //type:myInt kind:int64
	reflecType(c) //type:int32 kind:int32

	var e = map[int]string{
		01: "sss",
	}
	type person struct {
		name string
		age  int
	}
	type book struct {
		title string
	}
	p1 := person{
		name: "大壮",
		age:  12,
	}
	b1 := book{
		title: "这是八岁？",
	}
	reflecType(p1) //type:person kind:struct
	reflecType(b1) //type:book kind:struct

	fmt.Println(e)
	reflecType(e) //type: kind:map
	//Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空
}
