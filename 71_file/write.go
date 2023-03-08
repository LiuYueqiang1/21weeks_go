package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 文件写入操作
func write1() {
	fileObj, err := os.OpenFile("F:\\goland\\go_project\\21weeks_go\\71_file\\writetest.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0256)
	if err != nil {
		fmt.Printf("open file is failed,err:%v\n", err)
		return
	}
	defer fileObj.Close()
	str := "hello 沙河\n"
	fileObj.Write([]byte(str))
	fileObj.WriteString("hello 小王子\n")
}

// 按行写入
func bufio_NewWriter() {
	fileObi, err := os.OpenFile("F:\\goland\\go_project\\21weeks_go\\71_file\\writetest.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0234)
	if err != nil {
		fmt.Printf("open file is failed,err:%v\n", err)
		return
	}
	defer fileObi.Close()
	writer := bufio.NewWriter(fileObi)
	writer.WriteString("hello 沙河\n") //将数据先写入缓存
	writer.Flush()                   //将缓存中的内容写入文件
}

// 直接向文件中写入
func ioutilsritefile() {
	str := "hello 沙河筱往"
	err := ioutil.WriteFile("F:\\goland\\go_project\\21weeks_go\\71_file\\writetest.txt", []byte(str), 0253)
	if err != nil {
		fmt.Println("文件写入失败", err)
		return
	}
}
func main() {
	//write1()
	//bufio_NewWriter()
	ioutilsritefile()
}
