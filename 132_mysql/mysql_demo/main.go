package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

var db *sql.DB

func initDB() (err error) {
	//数据库信息
	dsn := "root:961024@tcp(localhost:3306)/db_student_manager_web"
	//连接数据库
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		//fmt.Printf("open %s failed,err:%v", dsn, err) //验证dsn格式是否正确
		return
	}
	err = db.Ping() //连接数据库
	if err != nil {
		//fmt.Printf("open %s failed,err:%v", dsn, err) //验证密码是否正确
		return
	}
	//fmt.Println("连接数据库成功")
	//设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	return
}

type user struct {
	id   int
	name string
	age  int
}

// 查询单条记录
func queryOne(id int) {
	var u1 user
	//1、写查询单条记录的sql语句
	sqlStr := `select id, name, age from user where id=?;`
	//2、执行
	//qdb.QueryRow(sqlStr, 2).Scan(&u1.id, &u1.name, &u1.age)
	roeObj := db.QueryRow(sqlStr, id) //从连接池里那一个连接出来去数据库查询单条记录
	//3、拿到结果
	roeObj.Scan(&u1.id, &u1.name, &u1.age) //Scan会释放数据库连接，归还到连接池中
	//打印结果
	fmt.Printf("u1:%#v\n", u1)
}
func insert() {

}
func main() {
	err := initDB()
	if err != nil {
		fmt.Println("init DB failed,err:", err)
	}
	fmt.Println("连接数据库成功")
	queryOne(2)
}
