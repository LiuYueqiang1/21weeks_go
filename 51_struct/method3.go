package main

import "fmt"

type person3 struct {
	name string
	age  int
}

func main() {
	var p person3
	p.name = "史莱姆"
	p.age = 100
	fmt.Println(p)
	var p1 = person3{
		name: "五条悟",
		age:  23,
	}

	fmt.Println(p1)
	//方法2
	s1 := []int{1, 2, 3, 4}
	m1 := map[string]int{
		"stu1": 100,
		"stu2": 20,
		"stu3": 50,
	}
	fmt.Println(s1, m1)

	p3 := person3{
		name: "维德鲁拉",
		age:  1000,
	}
	fmt.Println(p3)
}

//q3 为什么要有构造函数
func newPerson3(name string, age int) person3 {
	return person3{
		age:  age,
		name: name,
	}
}