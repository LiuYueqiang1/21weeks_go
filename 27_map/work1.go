package main

import "fmt"

// 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1
func main() {
	var str string = "how do you do"
	str2 := []rune(str)
	fmt.Println(len(str2))
	str3 := []rune{}
	//做一个map函数，出现相同的单词则value+1
	//	str4 := make(map[rune]int)
	for _, s := range str {
		if s == ' ' {
			continue
		}

		str3 = append(str3, s)

	}
	fmt.Printf("%c", str3)
	//fmt.Println(str4)
}
