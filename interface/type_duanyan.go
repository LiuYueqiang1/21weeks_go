package main

import "fmt"

// 类型断言
// 不仅可以输出类型，还可以调用这个
// 而%T不可以
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("字符串类型", t)
	case int:
		fmt.Println("整型", t)
	case int64:
		fmt.Println("int64类型", t)
	case bool:
		fmt.Println("bool型", t)
	}
}
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if ok {
		fmt.Println("字符串类型", str)
	} else {
		fmt.Println("不是字符串")
	}
}
func main() {
	assign2("你好")
}
