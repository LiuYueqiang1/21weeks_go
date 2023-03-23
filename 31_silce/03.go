package main

import "fmt"

func main() {
	var a = make([]string, 5, 10)
	var b = make([]int, 5, 10)
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	a[0] = "s"
	a[1] = "a"
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
