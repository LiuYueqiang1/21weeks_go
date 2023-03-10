package calc

import (
	"21weeks/21weeks_go/69_package/snow"
	"fmt"
)

// Name 是一个测试的全局变量，首字母需要大写
var Name = "沙河"

// Add 是把两个字母相加的函数
func Add(x, y int) (int, int) {
	snow.Snow()
	Sub(x, y)
	return x + y, Sub(x, y)
}

// init函数在导入的时候自动执行
// init函数没有参数也没有返回值
func init() {
	fmt.Println("calc.init()")
	fmt.Println(Name)
}
