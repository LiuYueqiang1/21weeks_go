package main

import (
	"fmt"
	"strings"
)

// 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1
// 按照空格划分成切片
// 将它们放到一个map中且统计出现的次数
func main() {
	var str string = "how do you do"
	s1 := strings.Split(str, " ")
	fmt.Println(s1)
	s2 := make(map[string]int)
	for _, v := range s1 {

		_, ok := s2[v]
		if !ok { //如果s2中没有s1的这些字符串，则=1
			s2[v] = 1
		} else { //如果有这些字符串，再＋1
			s2[v]++
		}
	}
	fmt.Println(s2)
	for i, v := range s2 {
		fmt.Println(i, v)
	}

}
