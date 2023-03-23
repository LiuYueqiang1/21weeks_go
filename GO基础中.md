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

#### 1、支持不同的地方输出日志

mylogger.go

##### 测试自己写的日志库

consloe.go：

```go
//向终端写日志

// 定义Logger结构体
type Logger struct {
}

// 给Logger建立构造函数去调用Logger
func Newlog() Logger {
   return Logger{}
}

// 给Logger定义一系列方法
func (l Logger) Debug(msg string) {
   fmt.Println(msg)
}
func (l Logger) Info(msg string) {
   fmt.Println(msg)
}
func (l Logger) Warning(msg string) {
   fmt.Println(msg)
}
func (l Logger) Error(msg string) {
   fmt.Println(msg)
}
func (l Logger) Fatal(msg string) {
   fmt.Println(msg)
}
```

Fprint：往指定位置写

##### 往终端中写

mylogger_test：是

```go
// 自定义一个日志库
func main() {
   log := mylogger.Newlog()
   for {
      log.Debug("这是一条Debug日志")
      log.Info("这是一条Info日志")
      log.Warning("这是一条Warning日志")
      log.Error("这是一条Erroe日志")
      log.Fatal("这是一条Fatal日志")
      time.Sleep(time.Second)
   }
}
```

可以在终端输出

##### 在console中将时间打印出来

```go
// 给Logger定义一系列方法
func (l Logger) Debug(msg string) {
   now := time.Now()
   TF := now.Format("2006-01-02 15:04:05")
   fmt.Printf("[%s] %s\n", TF, msg)
}
func (l Logger) Info(msg string) {
   now := time.Now()
   TF := now.Format("2006-01-02 15:04:05")
   fmt.Printf("[%s] %s\n", TF, msg)
}
func (l Logger) Warning(msg string) {
   now := time.Now()
   TF := now.Format("2006-01-02 15:04:05")
   fmt.Printf("[%s] %s\n", TF, msg)
}
func (l Logger) Error(msg string) {
   now := time.Now()
   TF := now.Format("2006-01-02 15:04:05")
   fmt.Printf("[%s] %s\n", TF, msg)
}
func (l Logger) Fatal(msg string) {
   now := time.Now()
   TF := now.Format("2006-01-02 15:04:05")
   fmt.Printf("[%s] %s\n", TF, msg)
}
```

#### 3、实现开关控制

日志要支持开关控制，比如说开发的时候什么级别都能输出，但是上线之后只有INFO级别往下的才能输出

```go
console.go
// 给log分级别
type LogLevel uint16 //这里必须是type，而不是var

const (
   DEBUG LogLevel = iota
   TRACE
   INFO
   WARNING
   ERROR
   FATAL
)

// 定义Logger结构体
type Logger struct {
   Lever LogLevel
}

// 给Logger建立构造函数去调用Logger,传入LogLevel
func Newlog(le LogLevel) Logger {
   return Logger{}
}

```

```
mylogger.go

log := mylogger.Newlog("Debug")
```

##### 建立构造函数调用Logger

但是主函数调用的时候Newlog输入的是String类型

所以需要将Newlog传入的改成这个

```go

// 给Logger建立构造函数去调用Logger
func Newlog(levelStr string) Logger {
   return Logger{}
}
```

修改完成后的：

```go
//将构造函数里面string类型的转换为Logger类型
func parseLogLevel(s string)(LogLevel,error){
   s=strings.ToLower(s)
   switch s {
   case "debug":
      return DEBUG,nil
   case "trace"
      return TRACE,nil
   case "info":
      return INFO,nil
   case "warning":
      return WARNING,nil
   case "error":
      return ERROR,nil
   case "fatal":
      return FATAL,nil
   default:
      err:=errors.New("无效的日志级别")
      return UNKNON , err
   }
}
// 给Logger建立构造函数去调用Logger
func Newlog(levelStr string) Logger {
   level,err:=parseLogLevel(levelStr)  //输入一个String类型的，返回一个Loglevel类型的
   if err!=nil{
      panic(err)
   }
   return Logger{
      Lever: level,
   }
}
```

将只要大于传入的日志级别输出

```go
func (l Logger) enable(loglevel LogLevel) bool {
   return loglevel >= l.Lever    //l.Lever：写入的
}

// 给Logger定义一系列方法
func (l Logger) Debug(msg string) {
   if l.enable(DEBUG) {
      now := time.Now()
      TF := now.Format("2006-01-02 15:04:05")
      fmt.Printf("[%s] [DEBUG] %s\n", TF, msg)
   }

}

func (l Logger) Info(msg string) {
   if l.enable(INFO) {
      now := time.Now()
      TF := now.Format("2006-01-02 15:04:05")
      fmt.Printf("[%s] [INFO] %s\n", TF, msg)
   }
}
func (l Logger) Warning(msg string) {
   if l.enable(WARNING) {
      now := time.Now()
      TF := now.Format("2006-01-02 15:04:05")
      fmt.Printf("[%s] [WARNING] %s\n", TF, msg)
   }
}
func (l Logger) Error(msg string) {
   if l.enable(ERROR) {
      now := time.Now()
      TF := now.Format("2006-01-02 15:04:05")
      fmt.Printf("[%s] [ERROR] %s\n", TF, msg)
   }

}
func (l Logger) Fatal(msg string) {
   if l.enable(FATAL) {
      now := time.Now()
      TF := now.Format("2006-01-02 15:04:05")
      fmt.Printf("[%s] [FATAL] %s\n", TF, msg)
   }
}
```

写出main()里面日志的文件名，函数名，行号

```go
// 执行的哪一行
// runtime 做程序运行时垃圾回收的操作，记录堆栈信息，函数调用，执行的什么文件
// pc调用的哪个函数
// file谁调用的这个函数
// 调用此函数的行号
func getInfo(skip int) (funcName, fileName string, lineNo int) {
   pc, file, lineNo, ok := runtime.Caller(skip)
   if !ok {
      fmt.Println("runtime.Caller() failed")
      return
   }
   funcName = runtime.FuncForPC(pc).Name()
   fileName = path.Base(file)
   return
}
```

将方法更新

```go
// 给Logger定义一系列方法
func (l Logger) Debug(msg string) {
   if l.enable(DEBUG) {
      funcName, fileName, lineNo := getInfo(2)
      now := time.Now()
      TF := now.Format("2006-01-02 15:04:05")
      fmt.Printf("[%s] [DEBUG] [%s:%s:%d] %s\n", TF, funcName, fileName, lineNo, msg)
   }
}
```

定义一个函数记录日志，将方法简化

```go
// 写一个记日志的函数
func log(lv LogLevel, msg string) {
   now := time.Now()
   TF := now.Format("2006-01-02 15:04:05")
   funcName, fileName, lineNo := getInfo(3)
   fmt.Printf("[%s] [DEBUG] [%s:%s:%d] %s\n", TF, funcName, fileName, lineNo, msg)
}
//此时的[DEBUG]需要在每个函数中声明
//传进来的是一个LogLevel类型，输出的应该是一个String类型

// 给Logger定义一系列方法
func (l Logger) Debug(msg string) {
   if l.enable(DEBUG) {
      log(DEBUG, msg)
   }
}

func (l Logger) Info(msg string) {
   if l.enable(INFO) {
      log(INFO, msg)
   }
}
func (l Logger) Warning(msg string) {
   if l.enable(WARNING) {
      log(WARNING, msg)
   }
}
func (l Logger) Error(msg string) {
   if l.enable(ERROR) {
      log(ERROR, msg)
   }

}
func (l Logger) Fatal(msg string) {
   if l.enable(FATAL) {
      log(FATAL, msg)
   }
}
```

将每个输出时的[错误]类型按照字符串输出，（原本是LogLevel类型）

```go
//F:\goland\go_project\21weeks\21weeks_go\82_test_rizhiku\mylogger\mylogger.go
func getLogString(lv LogLevel) string {
   switch lv {
   case DEBUG:
      return "DEBUG"
   case TRACE:
      return "TRACE"
   case INFO:
      return "INFO"
   case WARNING:
      return "WARNING"
   case ERROR:
      return "ERROR"
   case FATAL:
      return "FATAL"
   }
   return "DEBUG"
}
```

```go
// 修改记日志的函数
func log(lv LogLevel, msg string) {
   now := time.Now()
   TF := now.Format("2006-01-02 15:04:05")
   funcName, fileName, lineNo := getInfo(3)
   fmt.Printf("[%s] [%s] [文件名:%s 函数名:%s 行号:%d] %s\n", TF, getLogString(lv), fileName, funcName, lineNo, msg)
}
```

在main函数中输入的不只是String类型，输入的是任意类型

```go
//参考：
func Printf(format string, a ...any) (n int, err error) {
   return Fprintf(os.Stdout, format, a...)
}
//有格式化的语句，则传变量，没有格式化的语句，后面不传也可以
type any interface{}
```

后面的改动：

```go
// 写一个记日志的函数
func log(lv LogLevel, format string, a ...interface{}) {
   msg := fmt.Sprintf(format, a...) //Sprintf根据格式说明符格式化并返回结果字符串
   now := time.Now()
   TF := now.Format("2006-01-02 15:04:05")
   funcName, fileName, lineNo := getInfo(3)
   fmt.Printf("[%s] [%s] [文件名:%s 函数名:%s 行号:%d] %s\n", TF, getLogString(lv), fileName, funcName, lineNo, msg)
}
// 给Logger定义一系列方法
func (l Logger) Debug(format string, a ...interface{}) {
   if l.enable(DEBUG) {
      log(DEBUG, format, a...)
   }
}
func (l Logger) Info(format string, a ...interface{}) {
   if l.enable(INFO) {
      log(INFO, format, a...)
   }
}
func (l Logger) Warning(format string, a ...interface{}) {
   if l.enable(WARNING) {
      log(WARNING, format, a...)
   }
}
func (l Logger) Error(format string, a ...interface{}) {
   if l.enable(ERROR) {
      log(ERROR, format, a...)
   }

}
func (l Logger) Fatal(format string, a ...interface{}) {
   if l.enable(FATAL) {
      log(FATAL, format, a...)
   }
}
```

输出到文件中

将写日志的函数改为写日志的方法

```go
// 写一个记日志的方法
func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
   if c.enable(lv) {
      msg := fmt.Sprintf(format, a...) //Sprintf根据格式说明符格式化并返回结果字符串
      now := time.Now()
      TF := now.Format("2006-01-02 15:04:05")
      funcName, fileName, lineNo := getInfo(3)
      fmt.Printf("[%s] [%s] [文件名:%s 函数名:%s 行号:%d] %s\n", TF, getLogString(lv), fileName, funcName, lineNo, msg)
   }
}
```

新建一个file文件

```go
// 往文件里面写代码
type FileLogger struct {
   Level       LogLevel
   filePath    string
   fileName    string
   fileObj     *os.File
   errFileObj  *os.File
   maxFileSize int64
}
// FileLogger 的构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
   logLevel, err := parseLogLevel(levelStr)
   if err != nil {
      panic(err)
   }
   f1 := &FileLogger{
      Level:       logLevel,
      filePath:    fp,
      fileName:    fn,
      maxFileSize: maxSize,
   }
   err = f1.initFile() //按照文件路径和文件名将文件打开
   if err != nil {
      panic(err)
   }
   return f1
}
func (f *FileLogger) initFile() error {
   fullFileName := path.Join(f.filePath, f.fileName)
   fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
   if err != nil {
      fmt.Printf("open log file dailed,err:%v\n", err)
      return err
   }
   errfileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
   if err != nil {
      fmt.Printf("open err log file dailed,err:%v\n", err)
      return err
   }
   //日志文件已打开
   f.fileObj = fileObj
   f.errFileObj = errfileObj
   return nil
}
// 写一个记日志的方法
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
   if f.enable(lv) {
      msg := fmt.Sprintf(format, a...) //Sprintf根据格式说明符格式化并返回结果字符串
      now := time.Now()
      TF := now.Format("2006-01-02 15:04:05")
      funcName, fileName, lineNo := getInfo(3)
      fmt.Fprintf(f.fileObj, "[%s] [%s] [文件名:%s 函数名:%s 行号:%d] %s\n", TF, getLogString(lv), fileName, funcName, lineNo, msg)
      if lv >= ERROR {
         //如果要记录的日志大于等于ERROR级别，还需要再err日志中再记录一遍
         fmt.Fprintf(f.errFileObj, "[%s] [%s] [文件名:%s 函数名:%s 行号:%d] %s\n", TF, getLogString(lv), fileName, funcName, lineNo, msg)
      }
   }
}
// 写一个记日志的函数
func (f *FileLogger) enable(loglevel LogLevel) bool {
   return loglevel >= f.Level //f.Lever：写入的
}
// 给Logger定义一系列方法
func (f *FileLogger) Debug(format string, a ...interface{}) {
   f.log(DEBUG, format, a...)
}
func (f *FileLogger) Info(format string, a ...interface{}) {
   f.log(INFO, format, a...)
}
func (f *FileLogger) Warning(format string, a ...interface{}) {
   f.log(WARNING, format, a...)
}
func (f *FileLogger) Error(format string, a ...interface{}) {
   f.log(ERROR, format, a...)
}
func (f *FileLogger) Fatal(format string, a ...interface{}) {
   f.log(FATAL, format, a...)
}

func (f *FileLogger) Close() {
   f.fileObj.Close()
   f.errFileObj.Close()
}
```

在main函数中调用存入文件的方法

```go
//log := mylogger.Newlog("Debug")
log := mylogger.NewFileLogger("Info", "F:\\goland\\go_project\\21weeks\\21weeks_go\\82_test_rizhiku\\mylogger", "zhoulinwan.log", 10*1024*1024)
```

判断fileObj，errFileObj的大小是否超过了最大值，写一个方法

#### 5、根据文件大小分割文件

```go
// 切割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
   //2.备份一下 rename
   nowStr := time.Now().Format("20060102150405000")
   fileInfo, err := file.Stat()
   if err != nil {
      fmt.Printf("get file info failed,err:%v\n", err)
      return nil, err
   }
   logName := path.Join(f.filePath, fileInfo.Name())
   newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
   //1.关闭当前的日志文件
   file.Close()
   os.Rename(logName, newLogName)
   //3.打开一个新的日志文件
   fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
   if err != nil {
      fmt.Printf("open new log file failed, err:%v\n", err)
      return nil, err
   }
   //4.将打开的新日志文件对象赋值给 f.fileObj
   return fileObj, nil
}
```

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

![image-20230313084208141](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313084208141.png)

![image-20230313084906920](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313084906920.png)

![image-20230313085502384](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313085502384.png)

![image-20230313085548324](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313085548324.png)

![image-20230313085636466](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313085636466.png)

![image-20230313085654106](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313085654106.png)

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

## strconv标准库(类型转换)

把字符串解析成数据

string（）：拿着字符串做utf8编码找对应的符号 

```go
func ParseInt(s string, base int, bitSize int) (i int64, err error)
```

返回字符串表示的整数值，接受正负号。

base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；

bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。

```
func Sprintf(format string, a ...interface{}) string
```

Sprintf根据format参数生成格式化的字符串并返回该字符串。

```go
	str := "10000"
	//ret3 := int64(str)
	ret1, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("%#v %T\n", ret1, ret1) //10000 int64
	i := int32(97)
	ret4 := string(i)
	fmt.Println("ret4:", ret4)//a
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v", ret2) //"97"
```

ParsetInt：10进制，  

```
func Atoi(s string) (i int, err error)
```

Atoi是ParseInt(s, 10, 0)的简写。

```
func Itoa(i int) string
```

Itoa是FormatInt(i, 10) 的简写。

```
func ParseBool(str string) (value bool, err error)
```

返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。

```
func ParseFloat(s string, bitSize int) (f float64, err error)
```

解析一个表示浮点数的字符串并返回其值。

如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数（使用IEEE754规范舍入）。bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；返回值err是*NumErr类型的，语法有误的，err.Error=ErrSyntax；结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange

```go
func main() {
    //字符串转换为数字
   str := "10000"
   retInt, _ := strconv.Atoi(str)
    fmt.Printf("%#v %T\n", retInt, retInt)//10000 int
   fmt.Println(retInt) //10000
   //把数字转换为字符串类型
   i := int32(97)
   ret1 := string(i)
   fmt.Println(ret1) //a
   ret2 := fmt.Sprintf("%d", i)
   fmt.Printf("%#v\n", ret2) //"97"
   ret3 := strconv.Itoa(int(i))
   fmt.Printf("%#v\n", ret3) //"97"
   //从字符串中解析出bool值
   boolStr := "true"
   boolValue, _ := strconv.ParseBool(boolStr)
   fmt.Printf("%#v %T\n", boolValue, boolValue) //true bool
   //从字符串中解析出浮点数
   floatStr := "1.234"
   floatValue, _ := strconv.ParseFloat(floatStr, 64)
   fmt.Printf("%#v %T\n", floatValue, floatValue)//1.234 float64
}
```

## 并发***

### 引入

只是打印的比较慢，操作系统一直在输出

```go
func hello(i int) {
   fmt.Println("hello", i)
}

// 程序启动之后创建一个主goroutine去执行
func main() {
   for i := 0; i < 100; i++ {
      go hello(i) //开启一个单独的goroutine去执行hello函数
   }
   fmt.Println("main")
   time.Sleep(time.Second)
   //main函数结束了 由main函数启动的goroutine也结束了
}
```

闭包：可能会执行很多个gorountain

```go
func main() {
   for i := 0; i < 100; i++ {
      go func() {
         fmt.Println(i)
      }()
   }
   fmt.Println("main")
   time.Sleep(time.Second)
}

//执行了很多个匿名函数，因为一次循环的时间可以执行多个goroutine
//94
//9
//98
//99
//99
//99
//99
//100
//100
```

调用函数参数的i

```go
//给匿名函数传入循环的参数 
//这样的话一次只能调用一个
func main() {
   for i := 0; i < 100; i++ {
      go func(a int) {
         fmt.Println(a)
      }(i)
   }
   fmt.Println("main")
   time.Sleep(time.Second)
}
```

总结：启动go routine耗费一些资源和时间

如何使用优雅的方式等待go routine结束后再关闭主函数

**go routine什么时候结束？**

go routine对应的函数结束了，go routine就结束了

`main`函数执行完后，由`main`函数创建的那些go routine都结束了

### sync.WaitGroup  计数器

```go
var wg sync.WaitGroup

func f() {
   rand.Seed(time.Now().UnixNano())
   for i := 0; i < 5; i++ {
      r1 := rand.Int() //int64
      r2 := rand.Intn(10)
      fmt.Println(r1, r2)
   }
}
func f1(i int) {
	defer wg.Done() //计数器-1
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) //计数器+1
		//如果计数器变为零，则释放在Wait上阻塞的所有goroutine
		go f1(i)
	}
	wg.Wait() //等待直到counter变为0
}
```

**面试问题 goroutine调度模型：**

go routine和线程

go routine是用户态的线程，go语言自己实现的类似于线程的东西

M:N

goroutine初始栈的大小是2k

线程：操作系统的线程或者os线程

![image-20230313102826176](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313102826176.png)

![image-20230313103242895](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313103242895.png)

M：真正工作的

P：管理者

![image-20230313103556418](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313103556418.png)

![image-20230313104702643](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313104702643.png)

gomaxprocs（1）只有一个工作的   

gomaxprocs（6）默认cpu的逻辑核心数，默认跑满整个CPU

## channel

如果说 goroutine 是Go程序并发的执行体，`channel`就是它们之间的连接。`channel`是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制

```go
var b chan int//需要指定通道中元素的类型
```

通道的操作

1、发送：`ch1 <- 1`

2、接受：`x:=<- ch1`

`<- ch1`:丢了 

3、关闭：`close()`

### 无缓冲通道

```go
// 先从通道中获取数据，再向通道存入数据
var a []int
var b chan int
var wg sync.WaitGroup

func main() {
   fmt.Println(b) //<nil>
   b = make(chan int)
   wg.Add(1)
   go func() {
      defer wg.Done()
      x := <-b
      fmt.Println("后台goroutine从通道b中获取到了", x) //后台goroutine从通道b中获取到了 10
   }()
   b <- 10
   fmt.Println("10发送到通道b中了") //10发送到通道b中了
   b = make(chan int, 16)
   fmt.Println(b)//0xc00010a000
   wg.Wait()
}
```



```go
// 先向通道存入数据,再从通道中获取数据
var a1 []int
var b1 chan int
var wg1 sync.WaitGroup

func main() {
   fmt.Println(b1) //<nil>
   b1 = make(chan int)
   wg1.Add(1)
   go func() {
      defer wg1.Done()
      b1 <- 10
      fmt.Println("10发送到通道b中了", b1) //10发送到通道b中了 0xc00001c120
   }()

   x := <-b1
   fmt.Println("后台goroutine从通道b中获取到了", x) //后台goroutine从通道b中获取到了 10
   b1 = make(chan int, 16)
   fmt.Println(b1) //0xc00010a000
   wg1.Wait()
}
```

![channel-and-goroutines](https://img.draveness.me/2020-01-28-15802171487080-channel-and-goroutines.png)

上述两个 Goroutine，一个会向 Channel 中发送数据，另一个会从 Channel 中接收数据，它们两者能够**独立运行**并不存在直接关联，但是能通过 Channel 间接完成通信

### 有缓冲通道

**顺利执行**

```go
//有缓冲区的通道
var b2 chan int
func main() {
   fmt.Println(b2)
   b2 = make(chan int, 2)
   b2 <- 10
   fmt.Println("10发送到通道中了")
   b2 <- 20
   fmt.Println("20发送到通道中了")
   x := <-b2
   fmt.Println("从通道b中接收到了", x)
}
```

### chananl练习题

```go
var wg2 sync.WaitGroup

func f1(ch1 chan int) {
   defer wg2.Done()
   for i := 0; i < 100; i++ {
      ch1 <- i
   }
   close(ch1)
}
func f2(ch1, ch2 chan int) {
   defer wg2.Done()
   for {
      x, ok := <-ch1
      if !ok {
         break
      }
      ch2 <- x * x
   }
   close(ch2)
}
func main() {
   a := make(chan int, 100)
   b := make(chan int, 100)
   wg2.Add(2)
   go f1(a)
   go f2(a, b)
   wg2.Wait()
   for ret := range b {
      fmt.Println(ret)
   }
}
```

### 单向通道

多用于函数的参数里

```go
var wg2 sync.WaitGroup
var once sync.Once

func f1(ch1 chan <- int) { //只能向通道ch1里放值
   wg2.Done()
   for i := 0; i < 100; i++ {
      ch1 <- i
   }
   close(ch1)
}
func f2(ch1 <- chan int, ch2 chan <- int) { //只能从ch1里取值，只能向ch2放值
   wg2.Done()
   for {
      x, ok := <-ch1
      if !ok {
         break
      }
      ch2 <- x * x
   }
   once.Do(func() { close(ch2) }) //确保某个操作只执行一次
}

func main() {
   a := make(chan int, 100)
   b := make(chan int, 100)
   wg2.Add(3)
   go f1(a)
   go f2(a, b)
   go f2(a, b)
   wg2.Wait()
   for ret := range b {
      fmt.Println(ret)
   }
}
```

<img src="https://www.liwenzhou.com/images/Go/concurrence/channel.png" alt="img" style="zoom:120%;" />

nil：代表通道没有初始化

关闭成功后将数据读取完，继续读取将返回0值

三个goroutine开启5个通道

```go
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	//5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 5; a++ {
		<-results
	}
}
```

### 习题2

```go
//使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。
//开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
//开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
//主 goroutine 从resultChan取出结果并打印到终端输出

type job struct {
   value int64
}
type result struct {
   job *job
   sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func f1(z1 chan<- *job) {
   defer wg.Done()
   //循环生成int64类型的随机数，发送到jobChan中
   for {
      x := rand.Int63()
      newJob := &job{
         value: x,
      }
      z1 <- newJob
      time.Sleep(time.Millisecond * 500)
   }
}
func f2(z1 <-chan *job, resultChan chan<- *result) {
   defer wg.Done()
   for {
      job := <-z1
      sum := int64(0)
      n := job.value
      for n > 0 {
         sum += n % 10
         n = n / 10
      }
      newResult := &result{
         job: job,
         sum: sum,
      }
      resultChan <- newResult
   }
}
func main() {
   wg.Add(1)
   go f1(jobChan)
   wg.Add(24)
   for i := 0; i < 24; i++ {
      go f2(jobChan, resultChan)
   }
   for result := range resultChan {
      fmt.Printf("value %d  sum %d\n", result.job.value, result.sum)
   }
   wg.Wait()
}
```

### select多路复用

- 可处理一个或多个 channel 的发送/接收操作。
- 如果多个 case 同时满足，select 会**随机**选择一个执行。
- 对于没有 case 的 select 会一直阻塞，可用于阻塞 main 函数，防止退出。

- - - - - - - - - - - - - - - - - - -

### 互斥锁

#### 普通互斥锁

互斥锁是一种常用的控制共享资源访问的方法，它能够保证同一时间只有一个 goroutine 可以访问共享资源

```go
var x = 0
var wg sync.WaitGroup
var lock sync.Mutex //互斥锁
func add() {
   defer wg.Done()
   for i := 0; i < 10000; i++ {
      //修改之前加锁
      lock.Lock()
      x += 1
      //修改完成后解锁
      lock.Unlock()
   }
}
func main() {
   wg.Add(2)
   go add()
   go add()
   wg.Wait()
   fmt.Println(x)
}
```

**问题**：读网页和写网页的时候都加互斥锁将变得非常慢

#### 读写互斥锁

读写锁分为两种：读锁和写锁。当一个 goroutine 获取到读锁之后，其他的 goroutine 如果是获取读锁会继续获得锁，如果是获取写锁就会等待；而当一个 goroutine 获取写锁之后，其他的 goroutine 无论是获取读锁还是写锁都会等待。

**应用场景**：读的次数远远大于写的次数

```go
var wg1 sync.WaitGroup
var lo sync.Mutex
var rw sync.RWMutex
var y = 0

func read() {
   defer wg1.Done()
   //lo.Lock()
   rw.RLock()
   fmt.Println(y)
   time.Sleep(time.Millisecond)
   //lo.Unlock()
   rw.RUnlock()
}
func write() {
   defer wg1.Done()
   // lo.Lock()
   rw.RLock()
   y = y + 1
   time.Sleep(time.Millisecond * 5)
   //lo.Unlock()
   rw.RUnlock()
}
func main() {
   now := time.Now()
   for i := 0; i < 10; i++ {
      wg1.Add(1)
      go write()
   }
   time.Sleep(time.Second)
   for i := 0; i < 1000; i++ {
      wg1.Add(1)
      go read()
   }
   wg1.Wait()
   tsub := time.Now().Sub(now)
   fmt.Println(tsub)
}
```

## 进程、线程、协程

\* [进程(process):](#进程process)

  \* [定义](#定义)

  \* [特征](#特征)

  \* [进程状态：（三状态）](#进程状态三状态)

  \* [进程状态](#进程状态)

\* [线程：](#线程)

  \* [线程状态：](#线程状态)

\* [协程：](#协程)

\* [与线程的比较](#与线程的比较)

<!-- vim-markdown-toc -->

<!-- vim-markdown-toc GitLab -->

\* [进程(process):](#进程process)

  \* [定义](#定义)

  \* [特征](#特征)

  \* [进程状态：（三状态）](#进程状态三状态)

  \* [进程状态](#进程状态)

\* [线程：](#线程)

  \* [线程状态：](#线程状态)

\* [协程：](#协程)

\* [与线程的比较](#与线程的比较)

<!-- vim-markdown-toc -->

<!-- vim-markdown-toc Marked -->

\* [进程(process):](#进程(process):)

  \* [定义](#定义)

  \* [特征](#特征)

  \* [进程状态：（三状态）](#进程状态：（三状态）)

  \* [进程状态](#进程状态)

\* [线程：](#线程：)

  \* [线程状态：](#线程状态：)

\* [协程：](#协程：)

\* [与线程的比较](#与线程的比较)

<!-- vim-markdown-toc -->

<!-- vim-markdown-toc Redcarpet -->

\* [进程(process):](#进程-process)

  \* [定义](#定义)

  \* [特征](#特征)

  \* [进程状态：（三状态）](#进程状态：（三状态）)

  \* [进程状态](#进程状态)

\* [线程：](#线程：)

  \* [线程状态：](#线程状态：)

\* [协程：](#协程：)

\* [与线程的比较](#与线程的比较)

<!-- vim-markdown-toc -->

### \## 进程(process):

#### \### 定义

\- 狭义定义：进程就是一段程序的执行过程例如启动的某个app。

\- 广义定义：进程是一个具有独立功能的程序关于某个数据集合的一次运行活动。它是操作系统动态执行的基本单元，在传统的操作系统中,进程即是基本的分配单元，也是基本的执行单元。

#### \### 特征

1. 每个进程都有自己的地址空间，一般情况下，包含文本区域、数据区域、堆栈
2. 进程是执行中的程序，程序是一个没有生命的实体，只有处理器赋予程序生命时，它才能成为一个活动的实体，我们称之为进程
3. 进程本身不会运行，是线程的容器。线程不能单独执行，必须组成进程
4. 一个程序至少有一个进程，一个进程至少有一个线程
5. 对于操作系统来讲，一个任务就是一个进程，比如开一个浏览器就是启动一个浏览器进程。打开一款app就是打开一个进程，例如打开香哈就是运行了一个进程。
6. 有些进程还不止同时做一件事情。比如打开香哈，它可以同时进行看视频并且回复用户评论，在一个进程内部，要同时干多件事情。

#### \### 进程状态：（三状态）

1. 就绪：获取出CPU外的所有资源、只要处理器分配资源就可以马上执行
2. 运行：获得处理器分配的资源，程序开始执行
3. 阻塞：当程序条件不够的时候，需要等待提交满足的时候才能执行。

#### \### 进程状态

1. 创建状态：进程在创建时需要申请一个空白PCB，向其中填写控制和管理进程的信息，完成资源分配。如果创建工作无法完成，比如资源无法满足，就无法被调度运行，把此时进程所处状态称为创建状态
2. 就绪状态：进程已经准备好，已分配到所需资源，只要分配到CPU就能够立即运行
3. 执行状态：进程处于就绪状态被调度后，进程进入执行状态
4. 阻塞状态：正在执行的进程由于某些事件（I/O请求，申请缓存区失败）而暂时无法运行，进程受到阻塞。在满足请求时进入就绪状态等待系统调用
5. 终止状态：进程结束，或出现错误，或被系统终止，进入终止状态。无法再执行

### \## 线程：

1. 一个进程中至少有一个线程，不然就没有存在的意义
2. 在一个进程内部，要同时干多件事情，就需要同时运行多个子任务，我们把进程内的这些子任务叫做线程
3. 多线程就是为了同步完成多项任务(在单个程序中同时运行多个线程完成不同的任务和工作)，不是为了提高运行效率，而是为了提高资源使用效率来提高系统的效率
4. 一个简单的比喻，多线程就像是火车上的每节车厢，而进程就是火车。
5. 线程是程序执行流的最小单元。一个标准的线程由当前的线程ID、当前指令指针、寄存器和堆栈组成
6. 同一个进程中的多个线程之间可以并发执行

#### \### 线程状态：

\* 就绪：指线程具备运行的所有条件，逻辑上可以运行，在等待处理机

\* 运行：指线程占用处理机正在运行

\* 阻塞：线程在等待一个事件，逻辑上不可执行

\> 如果我们要同时执行多个任务怎么办?

启动多个进程，每个进程虽然只有一个线程，但是多个进程可以一块执行多个任务

启动一个进程，在一个进程内启动多个线程，这样多个线程也可以一块执行多个任务

#### 多任务：

### \## 协程：

\- 协程是一种用户态的**轻量级线程**，协程的调度完全由用户控制（进程和线程都是由cpu 内核进行调度）。协程拥有自己的寄存器上下文和栈。协程调度切换时，将寄存器上下文和栈保存到其他地方，在切回来的时候，恢复先前保存的寄存器上下文和栈，直接操作栈则基本没有内核切换的开销，可以不加锁的访问全局变量，所以上下文的切换非常快。

对于 进程、线程，都是有内核进行调度，有 CPU 时间片的概念，进行 抢占式调度（有多种调度算法）

\- 对于 协程(用户级线程)，这是对内核透明的，也就是系统并不知道有协程的存在，是完全由用户自己的程序进行调度的，因为是由用户程序自己控制，那么就很难像抢占式调度那样做到强制的 CPU 控制权切换到其他进程/线程，通常只能进行 协作式调度，需要协程自己主动把控制权转让出去之后，其他协程才能被执行到。

### goroutine 和协程区别

\- 本质上，goroutine 就是协程。 不同的是，Golang 在 runtime、系统调用等多方面对 goroutine 调度进行了封装和处理，当遇到长时间执行或者进行系统调用时，会主动把当前 goroutine 的CPU (P) 转让出去，让其他 goroutine 能被调度并执行，也就是 Golang 从语言层面支持了协程。Golang 的一大特色就是从语言层面原生支持协程，在函数或者方法前面加 go关键字就可创建一个协程。

### \## 与线程的比较

\+ 每个 goroutine (协程) 默认占用内存远比 Java 、C 的线程少。

goroutine：2KB（官方）

线程：8MB（参考网络）

\+ 线程和 goroutine 切换调度开销方面

线程/goroutine 切换开销方面，goroutine 远比线程小

线程：涉及模式切换(从用户态切换到内核态)、16个寄存器、PC、SP...等寄存器的刷新等。

goroutine：只有三个寄存器的值修改 - PC / SP / DX.

## 网络编程

![osi七层模型](https://www.liwenzhou.com/images/Go/socket/osi.png)

### 互联网分层协议

#### 物理层

我们的电脑要与外界互联网通信，需要先把电脑连接网络，我们可以用双绞线、光纤、无线电波等方式。这就叫做”实物理层”，它就是把电脑连接起来的物理手段。它主要规定了网络的一些电气特性，作用是负责传送0和1的电信号。

#### 数据链路层

单纯的0和1没有任何意义，所以我们使用者会为其赋予一些特定的含义，规定解读电信号的方式：例如：多少个电信号算一组？每个信号位有何意义？这就是”数据链接层”的功能，它在”物理层”的上方，确定了物理层传输的0和1的分组方式及代表的意义。早期的时候，每家公司都有自己的电信号分组方式。逐渐地，一种叫做”**以太网**”（Ethernet）的协议，占据了主导地位。

以太网规定，一组电信号构成一个数据包，叫做”帧”（Frame）。每一帧分成两个部分：标头（Head）和数据（Data）。其中”标头”包含数据包的一些说明项，比如发送者、接受者、数据类型等等；”数据”则是数据包的具体内容。”标头”的长度，固定为18字节。”数据”的长度，最短为46字节，最长为1500字节。因此，整个”帧”最短为64字节，最长为1518字节。如果数据很长，就必须分割成多个帧进行发送。

那么，发送者和接受者是如何标识呢？以太网规定，连入网络的所有设备都必须具有”网卡”接口。数据包必须是从一块网卡，传送到另一块网卡。网卡的地址，就是数据包的发送地址和接收地址，这叫做MAC地址。每块网卡出厂的时候，都有一个全世界独一无二的MAC地址，长度是48个二进制位，通常用12个十六进制数表示。前6个十六进制数是厂商编号，后6个是该厂商的网卡流水号。有了MAC地址，就可以定位网卡和数据包的路径了。

我们会通过ARP协议来获取接受方的MAC地址，有了MAC地址之后，如何把数据准确的发送给接收方呢？其实这里以太网采用了一种很”原始”的方式，它不是把数据包准确送到接收方，而是向本网络内所有计算机都发送，让每台计算机读取这个包的”标头”，找到接收方的MAC地址，然后与自身的MAC地址相比较，如果两者相同，就接受这个包，做进一步处理，否则就丢弃这个包。这种发送方式就叫做”广播”（broadcasting）。

#### 网络层

按照以太网协议的规则我们可以依靠MAC地址来向外发送数据。理论上依靠MAC地址，你电脑的网卡就可以找到身在世界另一个角落的某台电脑的网卡了，但是这种做法有一个重大缺陷就是以太网采用广播方式发送数据包，所有成员人手一”包”，不仅效率低，**而且发送的数据只能局限在发送者所在的子网络。也就是说如果两台计算机不在同一个子网络，广播是传不过去的**。这种设计是合理且必要的，因为如果互联网上每一台计算机都会收到互联网上收发的所有数据包，那是不现实的。

因此，必须找到一种方法区分**哪些MAC地址属于同一个子网络**，哪些不是。如果是同一个子网络，就采用广播方式发送，否则就采用”路由”方式发送。这就导致了”网络层”的诞生。它的作用是引进一套新的地址，使得我们能够区分不同的计算机是否属于同一个子网络。这套地址就叫做”网络地址”，简称”网址”。

“网络层”出现以后，每台计算机有了两种地址，一种是MAC地址，另一种是网络地址。两种地址之间没有任何联系，MAC地址是绑定在网卡上的，网络地址则是网络管理员分配的。网络地址帮助我们确定计算机所在的子网络，MAC地址则将数据包送到该子网络中的目标网卡。因此，从逻辑上可以推断，必定是先处理网络地址，然后再处理MAC地址。

规定网络地址的协议，叫做**IP协议**。它所定义的地址，就被称为IP地址。目前，广泛采用的是IP协议第四版，简称IPv4。IPv4这个版本规定，网络地址由32个二进制位组成，我们通常习惯用分成四段的十进制数表示IP地址，从0.0.0.0一直到255.255.255.255。

根据IP协议发送的数据，就叫做IP数据包。IP数据包也分为”标头”和”数据”两个部分：”标头”部分主要包括版本、长度、IP地址等信息，”数据”部分则是IP数据包的具体内容。**IP数据包的”标头”部分的长度为20到60字节，整个数据包的总长度最大为65535字节**。

#### 传输层

有了MAC地址和IP地址，我们已经可以在互联网上任意两台主机上建立通信。但问题是同一台主机上会有许多程序都需要用网络收发数据，比如QQ和浏览器这两个程序都需要连接互联网并收发数据，我们如何区分某个数据包到底是归哪个程序的呢？也就是说，我们还需要一个参数，表示这个数据包到底供哪个程序（进程）使用。这个参数就叫做”端口”（port），它其实是每一个使用网卡的程序的编号。每个数据包都发到主机的特定端口，所以不同的程序就能取到自己所需要的数据。

“端口”是0到65535之间的一个整数，正好16个二进制位。0到1023的端口被系统占用，用户只能选用大于1023的端口。有了IP和端口我们就能实现唯一确定互联网上一个程序，进而实现网络间的程序通信。

我们必须在数据包中加入端口信息，这就需要新的协议。最简单的实现叫做UDP协议，它的格式几乎就是在数据前面，加上端口号。UDP数据包，也是由”标头”和”数据”两部分组成：”标头”部分主要定义了发出端口和接收端口，”数据”部分就是具体的内容。UDP数据包非常简单，”标头”部分一共只有8个字节，总长度不超过65,535字节，正好放进一个IP数据包。

UDP协议的优点是比较简单，容易实现，但是缺点是可靠性较差，一旦数据包发出，无法知道对方是否收到。为了解决这个问题，提高网络可靠性，TCP协议就诞生了。TCP协议能够确保数据不会遗失。它的缺点是过程复杂、实现困难、消耗较多的资源。TCP数据包没有长度限制，理论上可以无限长，但是为了保证网络的效率，通常TCP数据包的长度不会超过IP数据包的长度，以确保单个TCP数据包不必再分割。

#### 应用层

应用程序收到”传输层”的数据，接下来就要对数据进行解包。由于互联网是开放架构，数据来源五花八门，必须事先规定好通信的数据格式，否则接收方根本无法获得真正发送的数据内容。”应用层”的作用就是规定应用程序使用的数据格式，例如我们TCP协议之上常见的Email、HTTP、FTP等协议，这些协议就组成了互联网协议的应用层。

如下图所示，发送方的HTTP数据经过互联网的传输过程中会依次添加各层协议的标头信息，接收方收到数据包之后再依次根据协议解包得到数据。

![HTTP数据传输图解](https://www.liwenzhou.com/images/Go/socket/httptcpip.png)

### socket编程

Socket是BSD UNIX的进程通信机制，通常也称作”套接字”，用于描述IP地址和端口，是一个通信链的句柄。Socket可以理解为TCP/IP网络的API，它定义了许多函数或例程，程序员可以用它们来开发TCP/IP网络上的应用程序。电脑上运行的应用程序通常通过”套接字”向网络发出请求或者应答网络请求。

#### socket图解

`Socket`是应用层与TCP/IP协议族通信的中间软件抽象层。在设计模式中，`Socket`其实就是一个门面模式，它把复杂的TCP/IP协议族隐藏在`Socket`后面，对用户来说只需要调用Socket规定的相关函数，让`Socket`去组织符合指定的协议数据然后进行通信。

![socket图解](https://www.liwenzhou.com/images/Go/socket/socket.png)

#### Go语言实现TCP通信

##### TCP协议

TCP/IP(Transmission Control Protocol/Internet Protocol) 即传输控制协议/网间协议，是一种面向连接（连接导向）的、可靠的、基于字节流的传输层（Transport layer）通信协议，因为是面向连接的协议，数据像水流一样传输，会存在黏包问题。

##### TCP服务端

一个TCP服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝网。因为Go语言中创建多个goroutine实现并发非常方便和高效，所以我们可以每建立一次链接就创建一个goroutine去处理。

TCP服务端程序的处理流程：

1. 监听端口
2. 接收客户端请求建立链接
3. 创建goroutine处理链接

我们使用Go语言的net包实现的TCP服务端代码如下：

```go
// tcp/server/main.go

// TCP server端

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}
```

将上面的代码保存之后编译成`server`或`server.exe`可执行文件

两个客户端之间通信：

```go
// 服务端
func main() {
   //1.本地端口启动服务
   listener, err := net.Listen("tcp", "127.0.0.1:2000")
   if err != nil {
      fmt.Println("strat tcp server on 127.0.0.1:2000 failed,err:", err)
      return
   }
   //2.等待客户端与我建立连接
   conn, err := listener.Accept()
   if err != nil {
      fmt.Println("accept failed,err:", err)
      return
   }
   //3.与客户端通信
   var tmp [128]byte
   n, err := conn.Read(tmp[:])
   if err != nil {
      fmt.Println("read from conn failed,err:", err)
      return
   }
   fmt.Println(string(tmp[:n]))
}
```

```go
// 客户端
func main() {
   //1. 与server端建立连接
   conn, err := net.Dial("tcp", "127.0.0.1:2000")
   if err != nil {
      fmt.Println("dial \"tcp\",\"127.0.0.1:2000\" failed,err:", err)
      return
   }
   //2.发送数据
   conn.Write([]byte("hello limuru!"))
   conn.Close()
}
```

多个客户端之间建立连接：

main：//2 for

//3 给一个goroutine 

server：

```go
func processConn(conn net.Conn) {
   //3.与客户端通信
   // var tmp [128]byte
   var tymp = make([]byte, 128)
   for {
      n, err := conn.Read(tymp)
      if err != nil {
         fmt.Println("read from conn failed,err:", err)
         return
      }
      fmt.Println(string(tymp[:n]))
   }

}

// 服务端
func main() {
   //1.本地端口启动服务
   listener, err := net.Listen("tcp", "127.0.0.1:2000")
   if err != nil {
      fmt.Println("strat tcp server on 127.0.0.1:2000 failed,err:", err)
      return
   }
   //2.等待客户端与我建立连接
   for {
      conn, err := listener.Accept()
      if err != nil {
         fmt.Println("accept failed,err:", err)
         return
      }
      go processConn(conn)
   }
}
```

client:

```go
func main() {
   //1. 与server端建立连接
   conn, err := net.Dial("tcp", "127.0.0.1:2000")//Dial 连接到指定网络上的地址
   if err != nil {
      fmt.Println("dial \"tcp\",\"127.0.0.1:2000\" failed,err:", err)
      return
   }
   //2.发送数据
   reader := bufio.NewReader(os.Stdin)
   for {
      fmt.Print("请说话：")
      msg, _ := reader.ReadString('\n')
      msg = strings.TrimSpace(msg)
      if msg == "exit" {
         break
      }
      conn.Write([]byte(msg))
   }
   conn.Close()
}
```

##### UDP协议

UDP协议（User Datagram Protocol）中文名称是用户数据报协议，是OSI（Open System Interconnection，开放式系统互联）参考模型中一种**无连接**的传输层协议，不需要建立连接就能直接进行数据发送和接收，属于不可靠的、没有时序的通信，但是UDP协议的实时性比较好，通常用于视频直播相关领域。

服务端

不需要 //2.等待客户端与我建立连接

```go
// UDP/server/main.go

// UDP server端
func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
```

客户端

```go
// UDP 客户端
func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	sendData := []byte("Hello server")
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败，err:", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
```

## 内容回顾

### 锁

`sync.Mutex`

是一个结构体，是值类型，给函数传参数的时候要传指针

### 两个方法

```go
var lock sync.Mutex
lock.Lock()
lock.UnLock()
```

### 为什么要用锁？

防止同一时刻多个goroutine操作同一个资源



### 互斥锁

### 读写互斥锁

#### 应用场景

适用于读多写少的场景

#### 特点

读的goroutine来了，获取的是读锁，后续的goroutine能读不能写

写的goroutine来了，获取的是写锁，后续的goroutine不管是读还是写都要等待获取锁

#### 使用

```go
var reLock sync.RWMutex
rwLock.RLock()//读
rwLock.RUnlock()

rwLock.Lock()//写
rwLock.Unlock()//
```



### 等待组

`sync.Waitgroup`

用来灯goroutine执行完再继续

是一个结构体,是值类型，**给函数传参的时候需要传指针**

#### 使用

```go
var sync.WaitGroup

wg.Add(1)
wg.Done()
wg.Wait()
```

### Sync.once

使用场景

某些场景只需要执行一次的时候就可以使用`sync.Once` 

```go
var once sync.Once
once.Do()
```

比如blog加载图片的例子

```go
//ch2是外部
f:=func(){
    close(ch2)
}//闭包，一个函数包含外部函数的引用，则是一个闭包
once.Do(f)
```

比如并发

```go

x,ok:=<- ch1//通道关闭时返回false
```

### sync.Map

#### 使用场景

并发操作一个map的时候，内置的map不是并发安全的

#### 使用

是一个开箱即用的并发安全的map

```go
var syncMap sync.Map
//Map[key] = value
syncMap.Store(key,value)
syncMap.Load(key)
syncMap.LoadOrStore()
syncMap.Delete()
syncMap.Range()
```

### 原子操作

Go语言内置了一些针对内置的基本数据类型的一些并发安全的操作

```go
var i int64=10
atomic.AddInt64(&i,1)
```

### 网络编程

#### 互联网协议

OSI七层模型

## http_server客户端

http_server

![image-20230322111700022](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230322111700022.png)

### 前端规则

HTTP：超文本传输协议

规定了浏览器和网站服务器之间的通信规则

HTML:超文本标记语言

**裸体的人**

学的就是标记的符号，标签

CSS：层叠样式表

**让人穿上衣服**/**化妆**

规定了HTML中标签的具体央视（颜色、背景、大小、位置、浮动）

JavaScript：一种跑在浏览器上的编程语言

**让人动起来**

## 单元测试122

### 切割字符串测试

![image-20230322162021300](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230322162021300.png)

![image-20230322162603669](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230322162603669.png)

![image-20230322162721911](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230322162721911.png)

### 测试组和子测试124

子测试：![image-20230322165706999](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230322165706999.png)

### 测试函数覆盖率

### 性能基准测试124

## flag包  

**获取命令行参数**

flag.type()

flag.string

flag.int

## 面试题

### 如何判断一个链表有没有闭环

![image-20230323092158679](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230323092158679.png)

## 内容回顾

net/http包的用法，如何发请求

![image-20230323093722101](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230323093722101.png)

![image-20230323093910156](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230323093910156.png)

![image-20230323094417814](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230323094417814.png)

![image-20230323094710373](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230323094710373.png)
