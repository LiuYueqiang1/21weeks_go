package main

import "fmt"

func main() {
	////获取用户输入
	//var s string
	//fmt.Scan(&s)
	//fmt.Println("输入的内容是：", s)

	//获取用户输入2
	var (
		name  string
		age   int
		class string
	)
//	fmt.Scanf("%s %d %s\n", &name, &age, &class)
//	fmt.Println(name, age, class)

	fmt.Scanln(&name, &age, &class)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, class)
}
}
