package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type user struct {
	ID   int
	Name string
	Age  int
}

func initDB() (err error) {
	dsn := "root:961024@tcp(localhost:3306)/db_student_manager_web"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

// 查询单条数据示例
func queryRowDemo() {
	sqlStr := `select id, name, age from user where id=?;`
	var u user
	err := db.Get(&u, sqlStr, 1) //单行查询
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

// 查询多行代码
func queryMultiRowDemo() {
	sqlStr := `select id ,name ,age from user where id > ?;`
	var users []user
	err := db.Select(&users, sqlStr, 0) //多行查询
	if err != nil {
		fmt.Println("query failed,err:", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

// DB.NamedExec方法用来绑定SQL语句与结构体或map中的同名字段
func insertUserDemo() (err error) {
	sqlStr := "INSERT INTO user (name,age) VALUES (:name,:age)"
	_, err = db.NamedExec(sqlStr,
		map[string]interface{}{
			"name": "七米",
			"age":  28,
		})
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
		return
	}
	//queryRowDemo()
	queryMultiRowDemo()
}
