package main

import (
	"encoding/json"
	"fmt"
)

// 首字母为什么要大写：格式化的功能是JSON包里的marshal方法里把p1所有东西拿出来转化成一个字符串
type person3 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person3{
		Name: "州立",
		Age:  20,
	}
	//JSON序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed,err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(b)) //{"name":"州立","age":20}
	//{"name":"州立","age":20}     用的JSON格式
	//{"Name":"州立","Age":20}	  未用JSON格式

	//JSON反序列化
	str := `{"name":"州立","age":20}`
	//var p2 person3
	//json.Unmarshal([]byte(str), &p2) //转化为字节类型的切片放入p2中
	//fmt.Printf("%v", p2)           //{州立 20}
	p2 := &person3{}
	json.Unmarshal([]byte(str), p2) //转化为字节类型的切片放入p2中
	fmt.Printf("%v", *p2)           //{州立 20}
}
