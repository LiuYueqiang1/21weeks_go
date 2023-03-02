package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b)) //已经执行了，只是给他放到一边而已
	fmt.Println(111)
	a = 0
	defer calc("2", a, calc("20", a, b))
	fmt.Println(222)
	b = 1
}

//1.a=1,b=2
//2.calc(1,1,calc("10",1,2))-->calc("1",1,3)   "10",1,2,3   //执行内部函数calc，先执行
//3. "1",1,3,4  //最后执行
//4. a=0
//5. calc("2",0,calc("20",0,2))-->calc("2",0,2)  "20",0,2,2   //执行内部函数calc，先执行
//6. "2",0,2,2  //倒数第二

//10 1 2 3
//111
//20 0 2 2
//222
//2 0 2 2
//1 1 3 4
