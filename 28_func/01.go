package main

import "fmt"

func f1(x int, y int) (ret int) { //ret 定义了return的返回值
	ret = x + y
	return //故在此处直接写个int
}
func f2(x int, y int) int { //参数必须要命名
	ret := x + y
	return ret //在函数中只命名了返回值类型，没有写返回值名称，故在return时要写上返回的是什么
}
func f3(x string, y ...int) { //y是可变长参数,切片类型的
	fmt.Println(x, y)

}
func main() {
	fmt.Println(f1(2, 3))
	f3("ss", 1, 2, 3, 5, 6, 4, 5, 6) //ss [1 2 3 5 6 4 5 6]
}
