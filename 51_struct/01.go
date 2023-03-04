package main

import "fmt"

type person struct {
	name, gender string
	age          int
}

// Go语言中函数参数永远是拷贝
func f(x person) {
	x.gender = "女"
}

// 传入指针
func f2(x *person) { //x *person  是person类型的指针，指向内存地址
	//(*x).gender = "女"     可以简写为下面的类型
	x.gender = "女" //语法糖，自动根据指针找到对应的变量
}

func main() {
	var p person
	p.name = "大风"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender) //男
	f2(&p)                // 传入到f2中的必须是地址
	fmt.Println(p.gender) //女
	//******************
	// 用new关键字对结构体进行实例化，得到的是结构体的地址
	var p2 = new(person)
	fmt.Printf("%T\n", p2)  //*main.person   类型
	fmt.Printf("%#v\n", p2) //&main.person{name:"", gender:""}    //是什么
	p2.age = 18
	p2.name = "打算"
	p2.gender = "男"
	fmt.Printf("%#v\n", p2) //&main.person{name:"打算", gender:"男", age:18}
	fmt.Printf("%T\n", p2)  //*main.person
	fmt.Printf("%p\n", p2)  //0xc000084180   p2保存的就是一块指向它保存数值的内存地址
	fmt.Printf("%p\n", &p2) //0xc0000ba020   &p2，p2的内存地址，&p2保存的数值只是一串数字
}
