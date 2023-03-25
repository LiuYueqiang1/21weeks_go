package main

import "fmt"

// 回文检测
func isPalindrome(a string) bool {

	//b := make(map[int]any, 10)
	//for in, va := range a {
	//	b[in] = va
	//}
	//fmt.Println(b)

	for i := 0; i < len(a); {
	tag:
		i++
		if a[i] == a[len(a)-i-1] {
			if i != len(a) {
				goto tag

			} else {
				return true
			}
		} else {
			return false
		}
	}
	return false
}
func main() {
	b := isPalindrome("油灯少灯油")
	fmt.Println(b)
}
