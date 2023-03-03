package main

import "fmt"

//递归:函数自己调用自己
//递归要有一个确定的出口，否则死循环
//递归适合处理那种问题相同\问题越来越小的场景
//3!=3*2*1
//4!=4*3*2*1

// 计算n的阶乘
func f(n uint64) uint64 { //u代表正数
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}
func main() {
	ret := f(7)
	fmt.Println(ret)
}
