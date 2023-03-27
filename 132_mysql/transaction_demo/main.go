package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

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

func transactionDemo() {
	//1、开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("开启事务失败,err::", err)
		return
	}
	sqlStr1 := `update user set age=age+10 where id =1;`
	sqlStr2 := `update user set age=age-10 where id =2;`
	ret1, err := tx.Exec(sqlStr1)
	if err != nil {
		//回滚
		tx.Rollback()
		fmt.Println("执行sql1出错，回滚事务,err:", err)
		return
	}
	n1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("get id failed,err:%v\n", err) //检测调用函数是否失败
		return
	}
	ret2, err := tx.Exec(sqlStr2)
	if err != nil {
		//回滚
		tx.Rollback()
		fmt.Println("执行sql2出错，回滚事务,err:", err)
		return
	}
	n2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("get id failed,err:%v\n", err) //检测调用函数是否失败
		return
	}
	fmt.Printf("sql1影响的行:%d，sql2影响的行:%d\n", n1, n2)

	//如果上面两次sql语句都执行成功，则提交事务
	err = tx.Commit()
	if err != nil {
		fmt.Println("提交事务失败！,err:", err)
		return
	}
	fmt.Println("事务执行成功！")
}
func main() {
	initDB()
	transactionDemo()
}
