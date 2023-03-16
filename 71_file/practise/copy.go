package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(dstName, srcName string) (written int64, err error) {
	//以读的方式打开源文件
	src, err := os.Open(srcName) //*file 文件名，err
	if err != nil {
		fmt.Printf("open %s failed,err:%v", srcName, err)
		return
	}
	defer src.Close()
	//以写的方式创建目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0625) //*file 文件名，err
	if err != nil {
		fmt.Printf("open %s failed,err:%v", srcName, err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
func main() {
	_, err := copyFile("F:\\goland\\go_project\\21weeks_go\\71_file\\copytest.txt", "F:\\goland\\go_project\\21weeks_go\\71_file\\writetest.txt")
	if err != nil {
		fmt.Println("copy failed", err)
		return
	}
	fmt.Println("copy success!")
}
