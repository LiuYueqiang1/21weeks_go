package main

import "fmt"

func main() {
	var s1 map[string]int
	s1 = make(map[string]int, 10) //开辟内存空间后
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
	fmt.Println("s2:", s2)
	fmt.Printf("%p\n", &s2)
	s2["ss"] = 18
	fmt.Printf("%p\n", &s2) //开辟内存空间后，里面的值怎么改变都没有影响，地址不会变
	fmt.Println("s1:", s1)

	delete(s1, "rob") //删除rob
	fmt.Println(s1)
}
