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

```
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
