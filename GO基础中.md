# GO基础中

## 80Time

```go
// 时区
func f2() {
   now := time.Now()
   fmt.Println(now) //加载本地时间
   //写出明天的时间
   //按照指定格式去解析一个字符串类型的时间
   time.Parse("2006-01-02 15:04:05", "2023-03-11 09:02:03")
   //按照东八区的时区和格式解析一个字符串的时间
   //根据字符串加载时区
   loc, err := time.LoadLocation("Asia/Shanghai")
   if err != nil {
      fmt.Println("load is failded,err:", err)
      return
   }
   //按照指定时区解析时间
   timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2023-03-11 09:02:03", loc)
   if err != nil {
      fmt.Println("prase is failed,err:", err)
      return
   }
   fmt.Println(timeObj)
   ts := timeObj.Sub(now)
   fmt.Println("时间差", ts)
}
func main() {
   f2()
}
2023-03-10 09:20:51.0355184 +0800 CST m=+0.005683001
2023-03-11 09:02:03 +0800 CST
时间差 23h41m11.9644816s
```

## 81日志库简单实现

###  需求分析

1、支持向不同的地方输出日志

2、日志分级别

1. ​	Debug：调试
2. ​	Info：记录正常事件
3. ​	Warning：可以写但是警告
4. ​	Error：错误
5. ​	Fatal：错误
6. ​	Trace

3、日志要支持开关控制，比如说开发的时候什么级别都能输出，但是上线之后只有INFO级别往下的才能输出

4、完整的日志记录要有时间、行号、文件名、日志级别、日志信息

5、日志文件要切割

### 实现

开发时做日志定位，开发时打开，上线后关闭，将日志写到一个文件里面，而不是在终端

1、打开一个文件，写处它的执行格式

2、将日志写入文件中

![image-20230310093206679](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230310093206679.png)

go程序：

mylogger.go

测试自己写的日志库

consloe.go：

![image-20230310112709204](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230310112709204.png)

![image-20230310112732547](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230310112732547.png)

Fprint：往指定位置写

往终端中写

mylogger_test：是

![image-20230310112753309](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230310112753309.png)

## 87 反射

**了解原理即可**

1、JSON格式化，动态获取传进来的变量类型

![image-20230311105133853](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230311105133853.png)

![image-20230311113526319](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230311113526319.png)

读取加载到main文件中

![image-20230311113833960](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230311113833960.png)

![image-20230312212719385](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230312212719385.png)

![image-20230312213446828](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230312213446828.png) 

![image-20230312213915443](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230312213915443.png)

![image-20230312214730622](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230312214730622.png)

![image-20230312215002748](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230312215002748.png)

![image-20230312215329343](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230312215329343.png)

![image-20230312220713152](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230312220713152.png)

![image-20230312221056685](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230312221056685.png)

根据结构体名字找数据： structName

![image-20230312221551571](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230312221551571.png)



![image-20230312221826006](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230312221826006.png)

![image-20230312221930892](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230312221930892.png)

![image-20230312222618717](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230312222618717.png)

判断conf是否有空行

![image-20230312222659718](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230312222659718.png)

2.3.2

拿着mysqoconfig这个字符串把这个结构体取出

找到=的位置，且判断去取值的 是否是一个结构体

![image-20230312223244774](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230312223244774.png)

### 作业

conf.ini文件

![image-20230312211833933](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230312211833933.png)



## 内容回顾

'2006-01-02 15:04:05.000'

### 时间类型

time.Time：time.Now()

时间戳:

time.Now().Unix():1970.1.1到现在的秒数

time.Now().UnixNano():1970.1.1到现在的纳秒数

time包的Time类型

### 时间间隔类型

time.Duration:时间间隔类型

time.Second

### 时间操作

事件对象+/.一个时间间隔对象

### 时间格式化

2006-01-02 15:04:05.000

### 定时器

每隔一秒钟执行一次

### 解析字符串格式的时间（时区）

### 日志库

之前学的内容的整合****

### 反射

接口类型的底层变量食分为两部分：动态类型和动态值

反射的应用：`json`等数据解析\ORM等工具

### 反射的两个方法

`reflect.TypeOf()`

`reflect.ValueOf()`

