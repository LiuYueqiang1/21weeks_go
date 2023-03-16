package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//创建一个文件1
	fileObj, err := os.OpenFile("F:\\goland\\go_project\\21weeks\\21weeks_go\\71_file\\review\\write1.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer fileObj.Close()
	str := "hello world！\n"
	fileObj.Write([]byte(str))
	fileObj.WriteString("你好1 good!\n")
	//读文件2
	filename, err := os.Open("F:\\goland\\go_project\\21weeks\\21weeks_go\\71_file\\review\\write1.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer filename.Close()
	reader := bufio.NewReader(filename)
	for {
		readtmp, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(readtmp)
	}

}
