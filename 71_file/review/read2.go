package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//按行读

func main() {
	filename, err := os.Open("F:\\goland\\go_project\\21weeks\\21weeks_go\\71_file\\review\\read2.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer filename.Close()
	//按行读取文件

	reader := bufio.NewReader(filename)
	for {
		tmpline, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("err:", err)
		}
		fmt.Print(tmpline)
	}

}
