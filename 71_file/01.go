package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readfromRead() {
	fileObj, err := os.Open("F:\\goland\\go_project\\21weeks_go\\71_file\\01.go")
	if err != nil {
		fmt.Printf("open file failed,%v\n", err)
		return
	}
	defer fileObj.Close()
	//读文件
	//var tmp [128]byte
	//for {
	//	n, err := fileObj.Read(tmp[:])
	//	if err != nil {
	//		fmt.Printf("Read is failed,%v\n", err)
	//		return
	//	}
	//	fmt.Printf("读了%v个字节\n", n)
	//	fmt.Println(string(tmp[:]))
	//	if n < 128 {
	//		return
	//	}
	//}
	//读文件2
	var tmp = make([]byte, 128) //读取指定长度
	for {
		n, err := fileObj.Read(tmp)
		//func (f *File) Read(b []byte) (n int, err error)
		//Read是一个方法，*File是接收者，读文件*file，写入到b中，返回一个读到的数量n的和一个错误
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		if err != nil {
			fmt.Printf("Read is failed,%v\n", err)
			return
		}
		fmt.Printf("读了%v个字节\n", n)
		fmt.Println(string(tmp))
		if n < 128 { //如果最后读的字节数<128，则这次已经读完了，直接return for循环
			return
		}
	}
}

// 读取文件方法2
func readfrombufio() {
	fileObj, err := os.Open("F:\\goland\\go_project\\21weeks_go\\71_file\\01.go")
	if err != nil {
		fmt.Printf("文件读取错误，为%v\n", err)
		return
	}
	defer fileObj.Close()
	//按行读取文件
	//创建一个用来从文件中读取内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n') //delim，数据源的意思，也就是说遇到换行就返回一个line,err
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			return
		}
		if err != nil {
			fmt.Printf("文件读取错误，%v\n", err)
			return
		}
		fmt.Print(line) //按行打印读到的
	}
}

// 读取文件方法3
// 直接读取文件
func readFromFileByIouttil() {
	ret, err := ioutil.ReadFile("F:\\goland\\go_project\\21weeks_go\\71_file\\01.go")
	if err != nil {
		fmt.Printf("文件读取错误，%v\n", err)
		return
	}
	fmt.Println(string(ret))
}
func main() {
	readFromFileByIouttil()
}
