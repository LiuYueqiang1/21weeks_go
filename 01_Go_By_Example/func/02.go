package main

import "fmt"

// 变参函数
// 这个函数使用任意数目的 int 作为参数 累和
func sum(a, b, c int) int {
	total := 0
	total = a + b + c
	return total
}
func main() {
	a := 10
	b := 2
	c := 30
	ret := sum(a, b, c)
	fmt.Println(ret)

}
