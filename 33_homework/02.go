package main

import (
	"fmt"
	"strings"
)

// 判断how do you do中单词出现的次数
func main() {
	s1 := "how do you do"
	//1、遍历这个句子
	fmt.Println(s1)
	s2 := strings.Split(s1, " ") //按照空格切割得到切片
	fmt.Println(s2)
	//2、按照空格将这个句子切割得到切片
	s3 := make(map[string]int)
	for _, v := range s2 {
		//fmt.Println(v)
		//_, ok := s3[v]
		//if !ok {
		//	s3[v] = 1
		//} else {
		//	s3[v]++
		//}
		//如果s3中存在v的话，则返回 值和ok，我们不需要值，只要ok，故用_代替
		if _, ok := s3[v]; !ok {
			//3、将遍历得到的切片放到一个map中
			s3[v] = 1
		} else {
			s3[v]++
		}
	}
	//4、累加出现的次数
	for k, v := range s3 {
		fmt.Println(k, v)
	}

}
