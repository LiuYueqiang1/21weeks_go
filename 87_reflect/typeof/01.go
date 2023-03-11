package main

import (
	"fmt"
	"reflect"
)

//使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type）

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Println(v)
}
func main() {
	var a float64 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)
}
