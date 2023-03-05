package main

import "fmt"

// 这里是一个函数，接受两个 int 并且以 int 返回它们的和
// Go 需要明确的返回值，例如，它不会自动返回最后一个表达式的值Go 需要明确的返回值，
// 例如，它不会自动返回最后一个表达式的值
func plus(a int, b int) int {

	c := a + b
	return c
}
func main() {
	var a int
	var b int

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	ret := plus(a, b)
	fmt.Println("a+b:", ret)
}
