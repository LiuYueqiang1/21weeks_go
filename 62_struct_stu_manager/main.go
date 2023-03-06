package main

import (
	"fmt"
	"os"
)

var mgr manStudent

func printmeau() {
	fmt.Println("利姆露大人您好！欢迎来到特佩斯特大王国管理系统")
	fmt.Println("请输入您的操作:")
	fmt.Println(`
1、查看所有魔物信息
2、添加新的魔物
3、修改魔物信息
4、将魔物从王国放逐
5、离开
`)
}
func main() {
	mgr = manStudent{
		allStudent: make(map[int]student, 100),
	}
	for {
		var choice int
		printmeau()

		fmt.Scanln(&choice)
		switch choice {
		case 1:
			mgr.readMu()
		case 2:
			mgr.addMu()
		case 3:
			mgr.editMu()
		case 4:
			mgr.deleMu()
		case 5:
			fmt.Println("利姆露大人，再见！")
			os.Exit(1)
		default:
			fmt.Println("利姆露大人，您的操作无效，请重新输入！")

		}
	}
}
