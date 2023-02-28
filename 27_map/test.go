package main

import "fmt"

// 遍历字符串
func traversalString() {
	s := "hello 世界"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
func main() {
	traversalString()
	// 在循环中使用append函数构建一个由九个rune字符构成的slice

	var r []rune
	for _, v := range "hello,世界" {
		r = append(r, v)
	}
	fmt.Printf("%q\n", r)
	fmt.Println(r)

	str := "hello 世界"
	str2 := []rune(str)
	fmt.Printf("%q\n", str2)
}
