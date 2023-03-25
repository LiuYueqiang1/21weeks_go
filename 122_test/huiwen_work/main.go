package main

import (
	"errors"
	"fmt"
)

// 回文检测
func Palindrome(input string) error {
	var arr = []rune(input)
	j := len(arr) - 1
	for i := 0; i < j; i++ {
		if arr[i] != arr[j] {
			var err = fmt.Sprintf("input[%v]:%v not eq to input[%v]:%v", i, j, string(arr[i]), string(arr[j]))
			return errors.New(err)
		}
		//fmt.Println(string(arr[i]))
		j--
	}
	return nil
}
func main() {
	b := Palindrome("油灯少灯油")
	fmt.Println(b)
}
