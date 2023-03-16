package main

import (
	"fmt"
	"io"
	"os"
)

// 打印内容
func main() {
	filename, err := os.Open("F:\\goland\\go_project\\21weeks\\21weeks_go\\71_file\\review\\read1.go")
	if err != nil {
		fmt.Println(err)
	}
	defer filename.Close()
	tmp := make([]byte, 128)
	for {
		n, err := filename.Read(tmp)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
		fmt.Println(string(tmp))
		if n < 128 {
			return
		}
	}

}
