package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filebyte, err := ioutil.ReadFile("F:\\goland\\go_project\\21weeks\\21weeks_go\\71_file\\review\\read2.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(filebyte))
}
