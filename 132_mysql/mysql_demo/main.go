package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

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
	//db.QueryRow(sqlStr, 2).Scan(&u1.id, &u1.name, &u1.age)
	roeObj := db.QueryRow(sqlStr, id) //从连接池里那一个连接出来去数据库查询单条记录
	//3、拿到结果
	roeObj.Scan(&u1.id, &u1.name, &u1.age) //Scan会释放数据库连接，归还到连接池中
	//打印结果
	fmt.Printf("u1:%#v\n", u1)
}

// 查询多条记录
func queryMore(n int) {
	//1、SQL语句
	sqlStr := `select id ,name,age from user where id >?;`
	//2、执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s is failed,err:%v\n", sqlStr, err)
		return
	}
	//3、一定要关闭rows
	defer rows.Close()
	//4、循环取值
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age) //跟单条查询不一样，没有关闭连接的
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
		}
		fmt.Printf("id:%d name:%s age:%d\n", u1.id, u1.name, u1.age)
	}
}

// ------------------------插入、更新和删除操作都使用Exec方法---------------------------
// 插入数据
func insert() {
	//1、写sql语句
	sqlStr := `insert into user(name,age) values("周芷若",28)`
	//2、exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed,err:%#v\n", err) //检测是否插入失败
		return
	}
	// 如果是插入数据的操作，能够拿到插入数据的id,刚才插入数据的id
	id, err := ret.LastInsertId() //
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err) //检测调用函数是否失败
		return
	}
	fmt.Println("id:", id) //查看插入的id
}

// 更新数据
func updateRowDemo() {
	//1、写sql语句
	sqlStr := "update user set age=? where id =?"
	//2、exec
	ret, err := db.Exec(sqlStr, 18, 4)
	if err != nil {
		fmt.Printf("update age where id is failed!,err:%#v\n", err)
		return
	}
	//3
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected is failed,err:%v\n", err)
		return
	}
	fmt.Println("影响的行", n)
}

// 删除数据
func deleteRowDemo() {
	//1、写sql语句
	sqlStr := "delete from user where id =?"
	//2、exec
	ret, err := db.Exec(sqlStr, 4)
	if err != nil {
		fmt.Printf("update age where id is failed!,err:%#v\n", err)
		return
	}
	//3
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected is failed,err:%v\n", err)
		return
	}
	fmt.Println("影响的行", n)
	fmt.Println("删除成功！")
}

// 预处理插入多条数据
func prepareInsert() {
	//把SQL语句分成两部分，命令部分与数据部分
	sqlStr := `insert into user(name,age) values(?,?);`
	stmt, err := db.Prepare(sqlStr) //把SQL语句先发给MySQL预处理，命令部分
	if err != nil {
		fmt.Printf("prepare failed,err:%#v\n", err)
		return
	}
	defer stmt.Close()
	// 后续只需要传入数据即可
	var m = map[string]int{
		"青衣蝠王": 58,
		"白毛鹰王": 59,
		"赵敏":   17,
		"玄冥大佬": 55,
		"张三丰":  108,
	}
	for k, v := range m {
		stmt.Exec(k, v)
	}
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Println("init DB failed,err:", err)
	}
	fmt.Println("连接数据库成功")
	//queryOne(2)
	//insert()
	prepareInsert()
	queryMore(0)
	//updateRowDemo()
	//deleteRowDemo()
}
