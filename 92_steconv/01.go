package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "10000"
	//ret3 := int64(str)
	ret1, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("%#v %T\n", ret1, ret1) //10000 int64
	i := int32(97)
	ret4 := string(i) //a
	fmt.Println("ret4:", ret4)
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n", ret2) //"97"

	retInt, _ := strconv.Atoi(str)
	fmt.Println(retInt)
}
