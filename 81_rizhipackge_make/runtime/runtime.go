package main

import (
	"fmt"
	"path"
	"runtime"
)

func f() {
	pc, file, line, ok := runtime.Caller(1) //runtime：记录函数运行时的信息
	if !ok {
		fmt.Println("runtime.Caller failed")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)        //main.main
	fmt.Println(file)            //F:/goland/go_project/21weeks/21weeks_go/81_rizhipackge_make/runtime/runtime.go
	fmt.Println(path.Base(file)) //runtime.go
	fmt.Println(line)            //22
}
func main() {
	f()
}
