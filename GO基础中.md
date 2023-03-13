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

错误写法：

![image-20230313090301317](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313090301317.png)

string（）：拿着字符串做utf8编码找对应的符号 

![image-20230313090805361](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313090805361.png)

![image-20230313091002697](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313091002697.png)

ParsetInt：10进制，  

![image-20230313091417886](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313091417886.png)

![image-20230313091525408](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313091525408.png)

![image-20230313091723181](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313091723181.png)

## 并发***

只是打印的比较慢，操作系统一直在输出

![image-20230313094700702](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313094700702.png)

闭包：可能会执行很多个gorountain

![image-20230313095014758](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313095014758.png)

调用函数参数的i

![image-20230313095108290](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313095108290.png)

总结：启动go routine耗费一些资源和时间

如何使用优雅的方式等待go routine结束后再关闭主函数

**go routine什么时候结束？**

go routine对应的函数结束了，go routine就结束了

`main`函数执行完后，由`main`函数创建的那些go routine都结束了

![image-20230313102118022](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313102118022.png)

![image-20230313102150764](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313102150764.png)

![image-20230313102219813](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313102219813.png)

![image-20230313102334021](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230313102334021.png)

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

<!-- vim-markdown-toc GFM -->

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
4. 一个简单的比喻，多线程就像是火车上的每节车厢，而进程就是火车
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
