package main

import "fmt"

func test() {
	var a map[int]string
	a = make(map[int]string, 10)
	fmt.Println(a[0] == "") //true
}
