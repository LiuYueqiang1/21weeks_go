package main

import (
	"fmt"
	"os"
)

var allStudent map[int]*student //声明学生变量

type student struct {
	ID   int
	Name string
}

func newstudent(id int, name string) *student {
	return &student{
		ID:   id,
		Name: name,
	}
}

func showStudent() {
	for i, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", i, v.Name)
	}
}
func addStudent() {
	var (
		id   int
		name string
	)
	fmt.Println("请输入学生的学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入学生的姓名：")
	fmt.Scanln(&name)
	//造学生
	newStu := newstudent(id, name)
	//追加到map中
	allStudent[id] = newStu
}
func delteStudent() {
	var deleteID int
	fmt.Println("请输入要删除的学号：")
	fmt.Scanln(&deleteID)
	delete(allStudent, deleteID)
}
func main() {
	allStudent = make(map[int]*student, 50)
	for {
		fmt.Println("欢迎来到学生管理系统")
		fmt.Println("请输入您的操作：")
		fmt.Println("1、查看所有学生信息")
		fmt.Println("2、添加学生")
		fmt.Println("3、删除学生信息")
		fmt.Println("4、退出学生信息管理系统")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			showStudent()
		case 2:
			addStudent()
		case 3:
			delteStudent()
		case 4:
			fmt.Println("再见！")
			os.Exit(1)
		default:
			fmt.Println("无效输入！")
		}
	}
}
