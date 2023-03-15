package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{
		name: "州立",
		age:  20,
	}
	//JSON序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed,err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(b)) //{}
}
