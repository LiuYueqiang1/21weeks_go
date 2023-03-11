package main

import (
	"fmt"
	"os"
)

//结构体版学生信息管理系统

func main() {
	//给学生管理者的map申请空间
	//调用结构体
	mgr := mgrStudent{
		allStudent: make(map[int]student, 50), //结构体初始化
	}

	fmt.Println("欢迎来到学生信息管理系统")
	for {
		fmt.Println(`1、查看所有学生信息
2、增加学生
3、修改学生信息
4、删除学生
5、退出系统
`)
		var input int
		fmt.Println("请输入您的操作：")
		fmt.Scanln(&input)
		switch input {
		case 1:
			{
				mgr.readMeau()
			}
		case 2:
			{
				mgr.addMeau()
			}
		case 3:
			{

			}
		case 4:
			{
				mgr.deleMenu()
			}
		case 5:
			fmt.Println("系统已退出")
			os.Exit(1)
		default:
			fmt.Println("输入错误,请重新输入：")
		}
	}

}
