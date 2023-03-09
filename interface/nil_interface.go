package main

import "fmt"

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 10)
	m1["name"] = "shadowracket"
	m1["age"] = 9000
	m1["merried"] = false
	m1["hobby"] = []string{"唱", "跳", "rap"}
	fmt.Println(m1) //map[age:9000 hobby:[唱 跳 rap] merried:false name:shadowracket]
	show(false)     //type:bool value:false
	show(m1)        //type:map[string]interface {} value:map[age:9000 hobby:[唱 跳 rap] merried:false  name:shadowracket]
}
