package main

import "fmt"

// 回文判断
func main() {
	s1 := "黄山落叶松叶落山黄"
	r := make([]rune, 0, len(s1))
	for _, v := range s1 {
		r = append(r, v)
	}
	fmt.Println(r)
	for i := 0; i < len(r)/2; i++ {
		//r[0]==r[len(r)-1]
		//r[1]==r[len(r)-2]
		//...
		//r[i]==r[len(r)-1-i]
		if r[i] == r[len(r)-1-i] {
			fmt.Println("是回文")
			return
		} else {
			fmt.Println("不是回文")
			return
		}
	}
}
