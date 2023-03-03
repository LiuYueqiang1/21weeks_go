package main

import "fmt"

func f1() {
	s1 := []int{1, 2, 3}
	s2 := s1
	var s3 []int
	copy(s3, s1)
	fmt.Println(s2) //[1 2 3]
	s2[1] = 100
	fmt.Println(s1)        //[1 100 3]
	fmt.Println(s2)        //[1 100 3]
	fmt.Println(s3)        //[]  因为只是定义了s3但是没有分配内存，故为空
	fmt.Println(s3 == nil) //true
}

func f2() {
	s1 := []int{1, 2, 3}
	s2 := s1
	var s3 = make([]int, 0, 3)
	copy(s3, s1)
	fmt.Println(s2) //[1 2 3]
	s2[1] = 100
	fmt.Println(s1)        //[1 100 3]
	fmt.Println(s2)        //[1 100 3]
	fmt.Println(s3)        //[]  make()函数定义好了长度，但是copy不会自动扩容，因为len=0，故仍然是空的，但是已经分配了内存，不是nil
	fmt.Println(s3 == nil) //false
}
func f3() {
	s1 := []int{1, 2, 3}
	s2 := s1
	var s3 = make([]int, 3, 3)
	copy(s3, s1)
	fmt.Println(s2) //[1 2 3]
	s2[1] = 100
	fmt.Println(s1)        //[1 100 3]
	fmt.Println(s2)        //[1 100 3]
	fmt.Println(s3)        //[1,2,3]
	fmt.Println(s3 == nil) //false
}
func main() {
	f3()
}
