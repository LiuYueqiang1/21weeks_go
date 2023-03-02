package main

import "fmt"

func a() {
	fmt.Println(1)
}
func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("释放数据库连接")
		} //如果有错误的话   执行错误，且释放数据库连接
	}()
	panic("错误！！！")
	fmt.Println(2)
}
func c() {
	fmt.Println(3)
}
func main() {
	a()
	b()
	c()
}
