package main

import (
	"fmt"
	"reflect"
)

// 类型（Type）和种类（Kind）
type myInt int64

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("%v,%v\n", v.Name(), v.Kind())
}

func main() {
	var a *float32
	var b myInt
	var c rune
	reflectType(a)
	reflectType(b)
	reflectType(c)
	type person struct {
		name  string
		age   int
		speak string
	}
	var p1 = person{
		name:  "麻宝锅",
		age:   69,
		speak: "来骗，来偷袭，我一个69岁的老同志",
	}
	reflectType(p1) //person,struct
}
