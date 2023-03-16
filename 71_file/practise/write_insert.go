package main

import (
	"fmt"
	"io"
	"os"
)

func f1() {
	//打开要操作的文件
	fileObj, err := os.OpenFile("F:\\goland\\go_project\\21weeks_go\\71_file\\copytest.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	//不能在文件中间插入内容，所以借助一个临时文件
	tmpObj, err := os.OpenFile("F:\\goland\\go_project\\21weeks_go\\71_file\\copytest.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("creat file failed,err:", err)
		return
	}
	defer tmpObj.Close()
	// 读取源文件写入临时文件
	var ret [5]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Println("Read from file failed,err:", err)
		return
	}
	//写入临时文件
	tmpObj.Write(ret[:n])
	//写入要插入的内容
	var s []byte
	str2 := "insert"
	s = []byte(str2)
	tmpObj.Write(s)
	//再把源文件后面的内容写入到临时文件中
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			fmt.Println("读取完毕")
			return
		}
		if err != nil {
			fmt.Println("文件读取失败,err:", err)
			return
		}
		tmpObj.Write(x[:n])
	}
	fileObj.Close()
	tmpObj.Close()
	err2 := os.Rename("F:\\goland\\go_project\\21weeks_go\\71_file\\copytest.txt", "F:\\goland\\go_project\\21weeks_go\\71_file\\copytest.tmp")
	if err2 != nil {
		fmt.Println("导入失败")
		return
	}
}
func main() {
	f1()
}
