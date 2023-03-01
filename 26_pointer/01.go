package main

import "fmt"

func p1() {
	//1.&:取地址
	n := 18
	p := &n
	fmt.Println(p)        //0xc00001a088
	fmt.Printf("%T\n", p) //*int:：int类型的指针，指向内存地址
	//2.*:根据地址取值
	m := *p
	fmt.Println(m)        //18
	fmt.Printf("%T\n", m) //int
}
func p2() {
	var a *int
	*a = 100
	fmt.Println(*a) //panic: runtime error: invalid memory address or nil pointer dereference
	//必须分配内存
}
func p3() {
	var a = new(int)
	fmt.Println(a)  //0xc00009e058
	fmt.Println(*a) //0
	//Go里面的指针只能读不能修改，不能修改指针变量指向的地址
	*a = 100
	fmt.Println(a)  //0xc00009e058
	fmt.Println(*a) //100
}
func p4() {
	//Go里面的指针只能读不能修改，不能修改指针变量指向的地址
	addr := "沙河"
	addeP := &addr
	fmt.Println(addeP)        //0xc000044050
	fmt.Printf("%T\n", addeP) //*string
	addeV := *addeP           //根据内存地址取值
	fmt.Println(addeV)        //沙河
	fmt.Println(&addeV)       //0xc000044060
	addeV = "纳扎"
	fmt.Println(&addeV) //0xc000044060
}
func main() {
	p4()
}
