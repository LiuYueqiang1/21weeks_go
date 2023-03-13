package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "10000"
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt)
	fmt.Println(retInt) //10000
	//把数字转换为字符串类型
	i := int32(97)
	ret1 := string(i)
	fmt.Println(ret1) //a
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n", ret2) //"97"
	ret3 := strconv.Itoa(int(i))
	fmt.Printf("%#v\n", ret3) //"97"
	//从字符串中解析出bool值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue) //true bool
	//从字符串中解析出浮点数
	floatStr := "1.234"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue) //1.234 float64
}
