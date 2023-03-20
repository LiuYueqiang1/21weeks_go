package main

import (
	"fmt"
	"reflect"
)

//使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type）

func reflectTV(x interface{}) {
	v1 := reflect.TypeOf(x)
	v2 := reflect.ValueOf(x)
	fmt.Println(v1)
	fmt.Println(v2)
}
func main() {
	var a float64 = 3.14
	reflectTV(a)
}
