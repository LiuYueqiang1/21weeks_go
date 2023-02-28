package main

import "fmt"

//map类型的切片和切片类型的map

func main() {
	//map类型的切片
	var s1 = make([]map[string]int, 3) //切片长度为3
	s1[0] = make(map[string]int)
	s1[0]["ss"] = 10
	s1[0]["aa"] = 20
	fmt.Println("s1:", s1) //[map[aa:20 ss:10] map[] map[]] 切片长度为3，map为加入的

	var s3 = make([]map[string]int, 3)
	s3[0] = map[string]int{
		"qq": 5,
		"ww": 10,
		"rr": 2,
	}
	fmt.Println("s3:", s3)
	//切片类型的map
	var s2 = make(map[string][]int, 5) //map长度为5
	s2["北京"] = []int{1, 2, 3}
	s2["newyork"] = []int{2, 3, 5}
	fmt.Println("s2:", s2) //map[newyork:[2 3 5] 北京:[1 2 3]]

	var s4 = make(map[string][]int, 5) //map长度为5
	s4["london"] = make([]int, 2, 10)  //切片长度为2，容量为10
	s4["london"][0] = 10
	s4["london"][1] = 20
	fmt.Println("s4:", s4) //map[london:[10 20]]

}
