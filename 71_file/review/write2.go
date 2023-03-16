package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileObj, err := os.OpenFile("F:\\goland\\go_project\\21weeks\\21weeks_go\\71_file\\review\\write2.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)                    //writer是一个缓冲区
	lenwritetemp, err := writer.WriteString("hello 利姆露！") //向缓冲区中存入内容
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(lenwritetemp)
	writer.Flush() //将缓冲区文字写入文本

	//读文件
	obj, _ := ioutil.ReadFile("F:\\goland\\go_project\\21weeks\\21weeks_go\\71_file\\review\\write2.txt")
	fmt.Println(string(obj))
}
