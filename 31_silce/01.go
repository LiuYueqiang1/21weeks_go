package main

import "fmt"

func main() {
	var s1 []int
	fmt.Println(s1)        //[],nil
	fmt.Println(s1 == nil) //true
	s1 = []int{1, 2, 3}
	fmt.Println(s1) //[1 2 3]
	//make初始化 分配内存
	s2 := make([]bool, 2, 4)
	fmt.Println(s2)         //[false false]
	s3 := make([]int, 0, 4) //分配了内存，但是长度为0，容量为4
	fmt.Println(s3 == nil)  //false
}
