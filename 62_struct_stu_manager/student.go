package main

import "fmt"

type student struct {
	id   int
	name string
}

type manStudent struct { //跟普通的struct一样声明类型即可
	allStudent map[int]student
}

func newStudent(id int, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

// 声明方法
// 查看所有魔物信息
func (m manStudent) readMu() {
	for _, v := range m.allStudent { //v是student类型
		fmt.Println("学号：", v.id, "姓名：", v.name)
	}
}
func (m manStudent) addMu() {
	var (
		id   int
		name string
	)
	fmt.Println("请输入魔物编号:")
	fmt.Scanln(&id)
	fmt.Println("请输入魔物名字:")
	fmt.Scanln(&name)
	stu := newStudent(id, name)
	m.allStudent[id] = *stu

}
func (m manStudent) editMu() {
	//1、获取管理者输入的编号
	var muId int
	fmt.Println("请输入魔物编号:")
	fmt.Scanln(&muId)
	//2、显示该编号对应的魔物信息，如果没有则显示查无此人

	stuObj, ok := m.allStudent[muId] //stuObj是一个student 类型
	if !ok {
		fmt.Println("查无此人")
		return //**************
	}
	fmt.Printf("你修改的魔物信息如下：编号:%d 名字:%s\n", stuObj.id, stuObj.name)
	fmt.Println("请输入魔物的新名字：")
	var newname string
	fmt.Scanln(&newname)
	stuObj.name = newname //更新魔物名字
	m.allStudent[muId] = stuObj
}
func (m manStudent) deleMu() {

	fmt.Println("请输入放逐的魔物编号:")
	var ID int
	fmt.Scanln(&ID)
	delete(m.allStudent, ID)
	fmt.Println("已放逐，不得再回特佩斯特王国！")
}
