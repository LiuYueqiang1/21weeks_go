package main

import "fmt"

//指针
func zeroval(ival int) {
	ival = 0
}
func zeroptr(iptr *int) { //输入的是一个指向地址的指针，没有返回值
	*iptr = 0 //地址对应的数为0
}

func main() {
	a := 1
	fmt.Println(a)
	zeroval(a)
	fmt.Println(a)
	zeroptr(&a)
	fmt.Println(a)
	
}