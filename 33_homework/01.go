package main

import (
	"fmt"
	"unicode"
)

// 判断字符串中汉字的数量
// 难点是判断一个字符是汉字
func main() {
	//1.依次拿到字符串中的字符
	s1 := "Hello沙河纳扎"
	count := 0
	for _, v := range s1 {
		//2.判断当前这个字符是不是汉字
		if unicode.Is(unicode.Han, v) {
			//3.把汉字出现的次数累加得到最终结果
			count++
		}
	}
	fmt.Println(count)

}
