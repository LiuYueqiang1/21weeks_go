package main

import "fmt"

type person2 struct {
	name string
	age  int
}

func newperson(na string, a int) *person2 {
	return &person2{
		name: na,
		age:  a,
	}
}
func main() {
	p1 := newperson("那个", 18)
	p2 := newperson("今年", 20)
	fmt.Println(*p1) //*person2 类型的，故p1，p2是地址，需要用*p1取地址
	fmt.Println(*p2)
	fmt.Printf("%T\n", p1.name)  //string类型  ，因为这是直接改变person2来赋值的，但是person2是 struct类型 ，并不是指向地址
	fmt.Printf("%p\n", &p1.name) //0xc000008078
	fmt.Printf("%p\n", &p1.age)  //0xc000008088
	fmt.Printf("%p\n", &p2.name) //0xc000008090
	fmt.Printf("%p\n", &p2.age)  //0xc0000080a0
	//一块结构体占用一块连续的内存
	//*************这个没有申请内存空间，只是直接改变了  地址对应的值，所以地址也变化了，从地址层面操作的，不然函数只能赋值，
	//new 申请一个内存空间，直接从地址层面操作，直接改变地址对应的值    指针和值不是一一对应关系，指针对应的值可以改变
	//上面也是从地址层面操作，不过改变的是值对应的地址，不然函数只有赋值值的操作，无法改变值，指针和值是一一对应关系
}
