package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}
	for _, stu := range stus {
		m[stu.name] = &stu //map[大王八:0xc000008078 娜扎:0xc000008078 小王子:0xc000008078]  地址就是stus的
	}
	fmt.Println(m) //map[大王八:{大王八 9000} 娜扎:{娜扎 23} 小王子:{小王子 18}]
	for k, v := range m {
		fmt.Println(v) //都是：{大王八 9000}
		fmt.Println(k, "=>", v.name)
	}
}
