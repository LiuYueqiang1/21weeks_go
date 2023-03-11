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
