package main

import (
	"encoding/json"
	"fmt"
)

// 首字母为什么要大写：格式化的功能是JSON包里的marshal方法里把p1所有东西拿出来转化成一个字符串
type person2 struct {
	Name string
	Age  int
}

func main() {
	p1 := person2{
		Name: "州立",
		Age:  20,
	}
	//JSON序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed,err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(b)) //{"Name":"州立","Age":20}
}
