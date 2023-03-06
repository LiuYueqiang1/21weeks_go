package main

import (
	"fmt"
	"os"
)

//01的学生信息管理系统的map[id]存的是学生，这次只存学生姓名

type student2 struct {
	id   int
	name string
}

// 构造student2的函数
func newstudent2(id int, name string) *student2 {
	return &student2{
		id:   id,
		name: name,
	}
}

var allStudent2 map[int]string

func showStudent2() {
	for i, v := range allStudent2 {
		//fmt.Println("学号：", i)
		//fmt.Println("姓名：", v)
		fmt.Printf("学号：%d 姓名：%s\n", i, v)
	}
}
func addStudent2() {

	var (
		id   int
		name string
	)
	fmt.Println("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名：")
	fmt.Scanln(&name)
	//用构造函数添加一个新学生
	newStu := newstudent2(id, name)
	allStudent2[id] = newStu.name
}
func delteStudent2() {
	fmt.Println("请输入删除的学生学号：")
	var id int
	fmt.Scanln(&id)
	delete(allStudent2, id)
}
func main() {
	allStudent2 = make(map[int]string, 50)
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
			showStudent2()
		case 2:
			addStudent2()
		case 3:
			delteStudent2()
		case 4:
			fmt.Println("再见！")
			os.Exit(1)
		default:
			fmt.Println("无效输入！")
		}
	}
}
