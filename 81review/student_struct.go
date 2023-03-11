package main

import "fmt"

type student struct {
	Id   int
	Name string
}

// 创建一个student的构造函数
func newStudent(id int, name string) *student {
	return &student{
		Id:   id,
		Name: name,
	}
}

// 创建一个学生管理者的函数
type mgrStudent struct {
	allStudent map[int]student
	//	student    //方法结构体测试函数
}

// 声明一个map类型的，写入学生
var allStudent map[int]student

// 声明所有的方法
func (s mgrStudent) readMeau() {
	//s.student.Name
	for i, v := range s.allStudent { //v获取的是student
		fmt.Printf("ID:%d Name:%s\n", i, v.Name)
	}
}

// 声明增加学生的方法
func (s mgrStudent) addMeau() {
	var (
		id   int
		name string
	)
	fmt.Println("请输入学生的学号：")
	fmt.Scanln(&id)
breakTag1:
	{
		for _, v1 := range s.allStudent {
			if v1.Id == 0 {
				goto breakTag2
			} else {

				for {
					for _, v := range s.allStudent {
						if id == v.Id {

							fmt.Println("id已存在，请重新输入新学号：")
							fmt.Scanln(&id)
							goto breakTag1

						} else {
							goto breakTag2
						}
					}
				}
			}
		}

	}
breakTag2:
	{
		fmt.Println("请输入学生的姓名：")
		fmt.Scanln(&name)
		stu := newStudent(id, name) //增加一个新学生
		s.allStudent[id] = *stu
	}
}

// 声明修改学生的方法
func (s mgrStudent) editMenu() {
	var (
		id   int
		name string
	)
	fmt.Println("请输入要修改的学生学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入要修改学生姓名：")
	fmt.Scanln(&name)
	//s.allStudent[id].Name=name  //无法分配给 s.allStudent[id].Name ****这很奇怪，
	//只能解释为s.allStudent[id]是一个struct类型，调用一个新的struct类型里面的东西的时候要用一个新的变量去 索引
	//51_struct  03 说明了
	newname := s.allStudent[id] //s.allStudent[id]    是一个student类型的
	newname.Name = name
}

// 声明删除学生的方法
func (s mgrStudent) deleMenu() {
	var (
		id int
	)
	fmt.Println("输入要删除的学生ID：")
	fmt.Scanln(&id)
	delete(s.allStudent, id)
	fmt.Println("删除完成！")
}
