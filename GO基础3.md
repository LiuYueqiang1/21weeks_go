# 今日内容day10

## MySQL

### 数据库

![image-20230325211111582](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230325211111582.png)

如果频繁的操作数据库时，打开数据库建立连接再关闭非常耗时，提前在连接池里打开若干个数据库，需要在数据库中执行增删改查的操作时直接取过来，用完再放回去，类似于goroutine



SQLite、MySQL、SQL server、postgresql、Oracle

关系型数据库：

用表来存一类的数据.

表结构设计的三大范式：《漫画数据库》

### MySQL知识点

SQL语句

DDL：操作数据库

DML：表的增删改查

DCL：用户及权限

### **Go操作MySQL**

database/sql

原生支持连接池，是并发安全的。

这个标准库没有具体的实现，只是列出来一些需要第三南方库实现的具体内容   

database/sql

原生支持连接池，是并发安全的。

这个标准库没有具体的实现，只是列出了

### 存储引擎

MySQL支持插件式存储引擎

#### MylSAM

1、查询速度快

2、只支持表锁

3、不支持事务

#### InnoDB

1、整体速度快

2、支持表锁和行锁

3、支持事务

事务：把多个SQL操作当成一个整体

事务的特点：

ACID：

1、原子性：事务要么成功要么失败，没有中间状态

2、一致性：数据的完整性没有被破坏

3、隔离性：事物之间是互相隔离的

​		隔离的四个级别

4、持久性：事务操作的结果是不会丢失的

索引：

索引的原理：B树和B+树

索引的类型

索引的命中

分库分表

SQL注入

SQL慢查询优化

MySQL主从：

MySQL读写分离

## Redis

KV数据库

## NSQ

Go语言开发轻量级的消息队列

## 包的依赖管理go module

Go1.1之后官方出的依赖管理工具

# MySQL

## 连接数据库

```go
//数据库信息
//用户名：密码@tcp()/数据库名称
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
```

## 查询

### 单行查询

```go
//1、写查询单条记录的sql语句
sqlStr := `select id, name, age from user where id=?;` //表名   ？占位符 
//2、执行
//qdb.QueryRow(sqlStr, 2).Scan(&u1.id, &u1.name, &u1.age)
roeObj := db.QueryRow(sqlStr, id) //从连接池里那一个连接出来去数据库查询单条记录
//3、拿到结果
roeObj.Scan(&u1.id, &u1.name, &u1.age) //Scan会释放数据库连接，归还到连接池中
//打印结果
fmt.Printf("u1:%#v\n", u1)
```

### 多行查询

```go
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
```

## 插入

```go
//插入数据
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
   id, err := ret.LastInsertId()
   if err != nil {
      fmt.Printf("get id failed,err:%v\n", err)
      return
   }
   fmt.Println("id:", id) //查看插入的id
}
```

## 增删改查

```go
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
func main() {
   err := initDB()
   if err != nil {
      fmt.Println("init DB failed,err:", err)
   }
   fmt.Println("连接数据库成功")
   //queryOne(2)
   //insert()
   //queryMore(0)
   //updateRowDemo()
   deleteRowDemo()
}
```

## 预处理

```go
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
```

## 事务

事务的方法：

Go语言中使用以下三个方法实现MySQL中的事务操作。 开始事务

```go
func (db *DB) Begin() (*Tx, error)
```

提交事务

```go
func (tx *Tx) Commit() error
```

回滚事务

```go
func (tx *Tx) Rollback() error
```

下面的代码演示了一个简单的事务操作，该事物操作能够确保两次更新操作要么同时成功要么同时失败，不会存在中间状态。

```go
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
   sqlStr1 := `update user set age=age-10 where id =1;`
   sqlStr2 := `update user set age=age+10 where id =2;`
   _, err = tx.Exec(sqlStr1)
   if err != nil {
      //回滚
      tx.Rollback()
      fmt.Println("执行sql1出错，回滚事务,err:", err)
      return
   }
   _, err = tx.Exec(sqlStr2)
   if err != nil {
      //回滚
      tx.Rollback()
      fmt.Println("执行sql2出错，回滚事务,err:", err)
      return
   }
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
```

## sqlx使用



go get github.com/jmoiron/sqlx

```go
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
```

### sqlx注入

