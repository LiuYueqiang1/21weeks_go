package main

import (
	"fmt"
	"reflect"
)

// Valueof修改值
func typeValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
}
func typeValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	var a int64 = 100
	typeValue1(&a)
	fmt.Println(a) //100
	typeValue2(&a)
	fmt.Println(a) //200
}
