package main

import "fmt"

func main() {
	var s1 map[string]int
	s1 = make(map[string]int, 10) //开辟内存空间
	s1["rob"] = 16
	s1["job"] = 14
	s2 := map[string]int{ //初始化2
		"ss": 16,
		"aa": 20,
	}
	v, ok := s1["rob"]
	if !ok {
		fmt.Println("查无此人")
	} else {
		fmt.Println(v)
	}
	fmt.Println(s2)
	fmt.Println(s1)

	delete(s1, "rob") //删除rob
	fmt.Println(s1)
}
