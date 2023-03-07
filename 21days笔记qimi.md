## 数组

```bash
var 数组变量名 [元素数量]T
```

## 切片

```go
var name []T
```

### make函数构造切片：

make([]T, size, cap)

**切片**：指针，长度，容量

**两种初始化方法**：一个直接传值，一个分配内存

```
var s1 []int
fmt.Println(s1)        //[],nil
fmt.Println(s1 == nil) //true
s1 = []int{1, 2, 3}
fmt.Println(s1) //[1 2 3] 
//make初始化 分配内存
s2 := make([]bool, 2, 4)
fmt.Println(s2)         //[false false]
s3 := make([]int, 0, 4) //分配了内存，但是长度为0，容量为4
fmt.Println(s3 == nil)  //false
```

**go一定要申请内存**

### 切片的复制拷贝：

对一个切片的修改会影响另一个切片的内容

```go
func main() {
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1) //[100 0 0]
	fmt.Println(s2) //[100 0 0]
}
```

由于切片是引用类型，切片不存值，所以s1和s2都指向了同一块内存地址，修改s2的同时也会修改s1。

### 	Go语言内建的`copy()`函数

可以迅速地将一个切片的数据复制到另外一个切片空间中，`copy()`函数的使用格式如下：

```bash
copy(destSlice, srcSlice []T)
```

没有内存无法存值：

```go
func f1() {
   s1 := []int{1, 2, 3}
   s2 := s1
   var s3 []int
   copy(s3, s1)
   fmt.Println(s2) //[1 2 3]
   s2[1] = 100
   fmt.Println(s1)        //[1 100 3]
   fmt.Println(s2)        //[1 100 3]
   fmt.Println(s3)        //[]  因为只是定义了s3但是没有分配内存，故为空
   fmt.Println(s3 == nil) //true
}

func f2() {
   s1 := []int{1, 2, 3}
   s2 := s1
   var s3 = make([]int, 0, 3)
   copy(s3, s1)
   fmt.Println(s2) //[1 2 3]
   s2[1] = 100
   fmt.Println(s1)        //[1 100 3]
   fmt.Println(s2)        //[1 100 3]
   fmt.Println(s3)        //[]  make()函数定义好了长度，但是copy不会自动扩容，因为len=0，故仍然是空的，但是已经分配了内存，不是nil
   fmt.Println(s3 == nil) //false
}
func f3() {
   s1 := []int{1, 2, 3}
   s2 := s1
   var s3 = make([]int, 3, 3)
   copy(s3, s1)
   fmt.Println(s2) //[1 2 3]
   s2[1] = 100
   fmt.Println(s1)        //[1 100 3]
   fmt.Println(s2)        //[1 100 3]
   fmt.Println(s3)        //[1,2,3]  
   fmt.Println(s3 == nil) //false
}
```

make定义好了长度，而copy不会自动扩容，故上述两端代码输出一样

//s3[]

### append()

Go语言的内建函数`append()`可以为切片动态添加元素。

 可以一次添加一个元素，可以添加多个元素，也可以添加另一个切片中的元素（后面加…）。



### 切片删除元素：

Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。 代码如下：

```go
func main() {
	// 从切片中删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a = append(a[:2], a[3:]...)
	fmt.Println(a) //[30 31 33 34 35 36 37]
}
```

## map

```go
map[KeyType]ValueType
```

```go
make(map[KeyType]ValueType, [cap])
```

Go语言中有个判断map中键是否存在的特殊写法，格式如下:

```go
value, ok := map[key]
```

### delete()

使用`delete()`内建函数从map中删除一组键值对

```go
delete(map, key)
```

## 指针pointer

```go
//1.&:取地址
n := 18
p := &n
fmt.Println(p)        //0xc00001a088
fmt.Printf("%T\n", p) //*int:：int类型的指针，指向内存地址
//2.*:根据地址取值
m := *p
fmt.Println(m)        //18
fmt.Printf("%T\n", m) //int
```

type：*int：int类型的指针，指向内存地址，所以我们用\*取出

空指针：

```go
var a *int
*a = 100
fmt.Println(*a) //panic: runtime error: invalid memory address or nil pointer dereference
//必须分配内存
```

### new()和make()

![image-20230303132258768](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230303132258768.png)



1. 二者都是用来做内存分配的。
2. make只用于slice、map以及channel申请内存的，make函数返回的是对应的这三个类型的本身；
3. new很少用，一般用来给基本数据类型申请内存，string、int，返回的是对应类型的指针（*string、*int）
4. 申请内存空间后，里面的值如何改变，这个内存地址都不变

```go
func p3() {
   var a = new(int)
   fmt.Println(a)  //0xc00009e058
   fmt.Println(*a) //0
   //Go里面的指针只能读不能修改，不能修改指针变量指向的地址
   *a = 100
   fmt.Println(a)  //0xc00009e058
   fmt.Println(*a) //100
    *a = 200
	fmt.Println(a)  //0xc00009e058    //申请内存空间后，里面的值如何改变，这个内存空间的地址都不变
	fmt.Println(*a) //200
}
func p4() {
   //Go里面的指针只能读不能修改，不能修改指针变量指向的地址
   addr := "沙河"
   addeP := &addr
   fmt.Println(addeP)        //0xc000044050
   fmt.Printf("%T\n", addeP) //*string
   addeV := *addeP           //根据内存地址取值
   fmt.Println(addeV)        //沙河
   fmt.Println(&addeV)       //0xc000044060
   addeV = "纳扎"
   fmt.Println(&addeV) //0xc000044060
}
```

## 函数：

## defer:

Go语言中函数的return不是原子操作，在底层是分为两步来执行

第一步：返回值赋值

函数中如果存在defer，那么defer执行的时机实在第一步和第二步之间

第二步：真正的RET返回

```go
func f1() int {
	x := 5
	defer func() {
		x++   //修改的是x，不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++  //函数传参 改变的是函数中的副本
	}(x)
	return 5  //返回值= x = 5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())//6
	fmt.Println(f3())
	fmt.Println(f4())
}
```

多个defer语句按照先进后厨的顺序延迟执行。

```go
func calc(index string, a, b int) int {
   ret := a + b
   fmt.Println(index, a, b, ret)
   return ret
}
func main() {
   a := 1
   b := 2
   defer calc("1", a, calc("10", a, b)) //已经执行了，只是给他放到一边而已
   fmt.Println(111)
   a = 0
   defer calc("2", a, calc("20", a, b))
   fmt.Println(222)
   b = 1
}

//1.a=1,b=2
//2.calc(1,1,calc("10",1,2))-->calc("1",1,3)   "10",1,2,3   //执行内部函数calc，先执行
//3. "1",1,3,4  //最后执行
//4. a=0
//5. calc("2",0,calc("20",0,2))-->calc("2",0,2)  "20",0,2,2   //执行内部函数calc，先执行
//6. "2",0,2,2  //倒数第二

//10 1 2 3
//111
//20 0 2 2
//222
//2 0 2 2
//1 1 3 4 
```

## 闭包***

闭包是一个函数，这个函数包含了他外部作用域的一个变量

底层原理：

1.函数可以作为返回值

2.函数内部查找变量的顺序，先在自己内部找，找不到往外层找

让f1调用f3

```go
// 要求让f1()调用f3()
func f1(f func()) {
   fmt.Println("this is f1")
   f()
}
func f2(x, y int) {
   fmt.Println("this is f2")
   fmt.Println(x + y)
}

func f3(x, y int) func() {
   tmp := func() {
      fmt.Println(x + y)
   }
   return tmp
}

func main() {
   ret := f3(100, 200)
   f1(ret)
   //this is f1
   //300 
   //f2未执行，执行的是f3的内置函数
}
```

```go
func f1(f func()) {
   fmt.Println("this is f1")
   f()
}
func f2(x, y int) {
   fmt.Println("this is f2")
   fmt.Println(x + y)
}

// 传入 f2函数以及它的两个传入值 ，传出的是一个函数
func f3(f func(int, int), x, y int) func() {
   tep := func() {
      f(x, y) //设置一个匿名函数调用传入的函数并返回
   }
   return tep
}

func main() {
   ret := f3(f2, 100, 200) //把原来需要传递两个int类型的参数包装成一个不需要传参的函数
   f1(ret)
   //this is f1
   //this is f2
   //300
}
```

## panic/recover

1. `recover()`必须搭配`defer`使用。
2. `defer`一定要在可能引发`panic`的语句之前定义。

```go
func a() {
   fmt.Println(1)
}
func b() {
   defer func() {
      err := recover()
      if err != nil {
         fmt.Println(err)
         fmt.Println("释放数据库连接")
      }  //如果有错误的话   执行错误，且释放数据库连接
   }()
   panic("错误！！！")
   fmt.Println(2)
}
func c() {
   fmt.Println(3)
}
func main() {
   a()
   b()
   c()
}
//1
//错误！！！    
//释放数据库连接
//3
```

## 结构体

### 匿名结构体（多用于临时场景）

![image-20230305172734707](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230305172734707.png)

51 函数内部修改的是副本

### 结构体的定义

只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。

```go
type person struct {
	name, gender string
	age          int
}

// Go语言中函数参数永远是拷贝
func f(x person) {
	x.gender = "女"
}

// 传入指针
func f2(x *person) { //x *person  是person类型的指针，指向内存地址
	//(*x).gender = "女"     可以简写为下面的类型
	x.gender = "女" //语法糖，自动根据指针找到对应的变量
}

func main() {
	var p person
	p.name = "大风"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender) //男
	f2(&p)                // 传入到f2中的必须是地址
	fmt.Println(p.gender) //女
	//******************
	// 用new关键字对结构体进行实例化，得到的是结构体的地址
	var p2 = new(person)
	fmt.Printf("%T\n", p2)  //*main.person   类型
	fmt.Printf("%#v\n", p2) //&main.person{name:"", gender:""}    //是什么
	p2.age = 18
	p2.name = "打算"
	p2.gender = "男"
	fmt.Printf("%#v\n", p2) //&main.person{name:"打算", gender:"男", age:18}
}
```

p保存的值对应的地址，p的地址

![image-20230304104440549](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230304104440549.png)

结构体指针赋值方式

```go
fmt.Printf("%T\n", p2)  //*main.person
fmt.Printf("%p\n", p2)  //0xc000084180   p2保存的就是一块指向它保存数值的内存地址
fmt.Printf("%p\n", &p2) //0xc0000ba020   &p2，p2的内存地址，&p2保存的数值只是一串数字
```

一个结构体占用一块连续的内存

### 构造函数（包含地址的区别）

**方便直接使用结构体**

```go
type person2 struct {
   name string
   age  int
}
//构造函数                         返回值的类型
func newperson(na string, a int) *person2 {
   return &person2{
      name: na,
      age:  a,
   }
}
func main() {
   p1 := newperson("那个", 18)
   p2 := newperson("今年", 20)
   fmt.Println(*p1) //*person2 类型的，故p1，p2是地址，需要用*p1取地址
   fmt.Println(*p2)
   fmt.Printf("%T\n", p1.name)  //string类型  ，因为这是直接改变person2来赋值的，但是person2是 struct类型 ，并不是指向地址
   fmt.Printf("%p\n", &p1.name) //0xc000008078
   fmt.Printf("%p\n", &p1.age)  //0xc000008088
   fmt.Printf("%p\n", &p2.name) //0xc000008090
   fmt.Printf("%p\n", &p2.age)  //0xc0000080a0
   //一块结构体占用一块连续的内存
   //*************这个没有申请内存空间，只是直接改变了  地址对应的值，所以地址也变化了，从地址层面操作的，不然函数只能赋值，
  
    //new 申请一个内存空间，直接从地址层面操作，直接改变地址对应的值    指针和值不是一一对应关系，指针对应的值可以改变
   //上面也是从地址层面操作，不过改变的是值对应的地址，不然函数只有赋值值的操作，无法改变值，指针和值是一一对应关系
}
//参考看目录  pointer
```

结构体是值类型，赋值都是拷贝

### 方法和接收者

Go语言中的`方法（Method）`是一种作用于特定类型变量的函数。这种特定类型变量叫做`接收者（Receiver）`。接收者的概念就类似于其他语言中的`this`或者 `self`。

方法的定义格式如下：

```go
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}   
//方法是接收者的函数，接收者指的是   哪个类型的变量   可以调用这个函数
```

方法

当在方法里需要**修改结构体变量的值**时需要用**指针**接收者

例如：过年年龄+1的方法



```go
type dog struct {
   name string
}

// 构造函数  调用结构体里面的东西 返回的是结构体名称
func newDog(name string) *dog {
   return &dog{
      name: name,
   }
}
// 方法是作用于特定类型的函数
//
// func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//    函数体
// }
//
// 传入的   传出的
//接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
//只有接收者这个类型的变量可以调用这个函数
func (d dog) wang() {
   fmt.Printf("%s:汪汪汪\n", d.name)
}
func main() {
   d1 := newDog("zzz") //返回的是结构体里面的东西，给d1
   d1.wang()
}
```

**标识符：变量名 函数名 类型名 方法名**

go语言中如果标识符首字母是大写的，就表示对外部可见   

### 值接收者和指针接收者：

1. 需要修改接收者中的值
2. 接收者是拷贝代价比较大的大对象
3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
4. 一般都是用指针接收者

 只能给自己定义的类型添加方法

```go
//给自定义类型加方法
//不能给别的暴力的类型添加方法，只能给自己包里的类型添加方法

type newInt int

// 方法    接受变量   方法名
func (n newInt) hello() {
   fmt.Println("这是一个int类型的方法")
}

func main() {
   n1 := newInt(100)
   fmt.Println(n1)
   n1.hello()
}
```

结构体遇到的问题：

myInt（100）是个什么？

```go
//方法1：
var x int32
x=10
//方法2：
var x2 int32=10
//方法3
var x3=int32(10)
//方法4
x4:=int32(10)
fmt.Println(x,x2,x3,x4)
```

```go
//方法1
var n1 newInt
n1 = 100
//方法2
var n2 newInt = 100
//方法3
var n3 = newInt(100)
//方法4
n4 := newInt(100) //强制类型转换
n4.hello()   
fmt.Println(n1, n2, n3, n4)
```

初始化：

```go
type person3 struct {
   name string
   age  int
}

func main() {
   var p person3
   p.name = "史莱姆"
   p.age = 100
   fmt.Println(p)
   var p1 = person3{
      name: "五条悟",
      age:  23,
   }

   fmt.Println(p1)
   //方法2
   s1 := []int{1, 2, 3, 4}
   m1 := map[string]int{
      "stu1": 100,
      "stu2": 20,
      "stu3": 50,
   }
   fmt.Println(s1, m1)

   p3 := person3{
      name: "维德鲁拉",
      age:  1000,
   }
   fmt.Println(p3)
}
```

3.为什么要有构造函数：

```go
//q3 为什么要有构造函数
func newPerson3(name string, age int) person3 {
   return person3{
      age:  age,
      name: name,
   }
```

### 匿名字段

```go
// 匿名字段
// 字段比较少比较简单
// 不常用
type person4 struct {
   string
   int
}

func main() {
   p1 := person4{
      "米栗木",
      18,
   }
   fmt.Println(p1)
}
```

### 嵌套结构体

```go
type person5 struct {
   name string
   age  int
}

// 嵌套结构体
type company struct {
   name string
   pe   person5
}

func main() {
   fmt.Println(p1)
   c1 := company{
      name: "华强集团",
      pe: person5{
         "刘华强",
         28,
      },
   }
   fmt.Println(c1)   //{华强集团 {刘华强 28}}
}
```

#### 匿名嵌套结构体：***

用的比较多

```go
type person5 struct {
   name string
   age  int
}
type address struct {
   city string
   mail string
}
// 嵌套结构体
type company struct {
   name    string
   pe      person5
   address //匿名嵌套结构体
}

func main() {
   c1 := company{
      name: "华强集团",
      pe: person5{
         "刘华强",
         28,
      },
   }
   fmt.Println(c1) //{华强集团 {刘华强 28}}
   c2 := company{
      name: "撒日朗",
      pe: person5{
         name: "华强",
         age:  28,
      },
      address: address{
         "北京不知名水果摊",
         "保熟吗.com",
      },
   }
   fmt.Println(c2.pe.name) //普通嵌套结构体    //华强
   fmt.Println(c2.city)    //先在自己的结构体里查找该字段，找不到就去匿名嵌套的结构体中查找  //北京不知名水果摊
   fmt.Println(c2)         //{撒日朗 {华强 28} {北京不知名水果摊 保熟吗.com}}
}
```

### 继承*******

```go
// 继承
type animal struct {
   name string
}

func (a animal) move() {
   fmt.Printf("%s会动\n", a.name)
}

type dog struct {
   feet   uint8
   animal //animal拥有的方法 和结构体，此时狗也拥有了
}

func (d dog) wang() {
   fmt.Printf("%s会汪汪汪\n", d.name)
}
func main() {
   d1 := dog{
      feet: 4,
      animal: animal{
         name: "岚牙",
      }, //类似于匿名嵌套结构体，但这是继承，可以使用animal的结构体
   }
   fmt.Println(d1) //{4 {岚牙}}
   d1.wang()       //岚牙会汪汪汪
   d1.move()       //继承自animal的方法    //岚牙会动
}
```

### 自定义类型和类型别名

![image-20230305195431512](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230305195431512.png)

##  JSON

www.json.cn

1、把Go语言中结构体变量---->JSON格式的字符串             序列化

2、JSON格式的字符串---->Go语言中能够识别的结构体变量            反序列化

```go
**********************1************************
type person struct {
   name string
   age  int
}

func main() {
   p1 := person{
      name: "州立",
      age:  20,
   }
   //JSON序列化
   b, err := json.Marshal(p1)
   if err != nil {
      fmt.Printf("marshal failed,err:%v", err)
      return
   }
   fmt.Printf("%v\n", string(b)) //{}
}

**********************2************************
// 首字母为什么要大写：格式化的功能是JSON包里的marshal方法里把p1所有东西拿出来转化成一个字符串
type person2 struct {
	Name string
	Age  int
}

func main() {
	p1 := person2{
		Name: "州立",
		Age:  20,
	}
	//JSON序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed,err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(b)) //{"Name":"州立","Age":20}
}
```

```go
**********************3************************
// 首字母为什么要大写：格式化的功能是JSON包里的marshal方法里把p1所有东西拿出来转化成一个字符串
type person3 struct {
   Name string `json:"name"`
   Age  int    `json:"age"`
}

func main() {
   p1 := person3{
      Name: "州立",
      Age:  20,
   }
   //JSON序列化
   b, err := json.Marshal(p1)
   if err != nil {
      fmt.Printf("marshal failed,err:%v", err)
      return
   }
   fmt.Printf("%v\n", string(b))
   //{"name":"州立","age":20}     用的JSON格式
   //{"Name":"州立","Age":20}     未用JSON格式

   //JSON反序列化   反序列化时要传递指针
   str := `{"name":"州立","age":20}`
   //var p2 person3
   //json.Unmarshal([]byte(str), &p2) //转化为字节类型的切片放入p2中
   //fmt.Printf("%v", p2)           //{州立 20}
   p2 := &person3{}
   json.Unmarshal([]byte(str), p2) //转化为字节类型的切片放入p2中
   fmt.Printf("%v", *p2)           //{州立 20}
}
```

![image-20230305195232197](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230305195232197.png)

## 学生信息管理系统

### 函数版

写一个能够查看、新增、删除学生的系统

功能：

1、打印菜单

2、等待用户执行操作

3、执行对应的函数

```go
var allStudent map[int]*student //声明学生变量

type student struct {
   ID   int
   Name string
}

func newstudent(id int, name string) *student {
   return &student{
      ID:   id,
      Name: name,
   }
}

func showStudent() {
   for i, v := range allStudent {
      fmt.Printf("学号：%d 姓名：%s\n", i, v.Name)
   }
}
func addStudent() {
   var (
      id   int
      name string
   )
   fmt.Println("请输入学生的学号：")
   fmt.Scanln(&id)
   fmt.Println("请输入学生的姓名：")
   fmt.Scanln(&name)
   //造学生
   newStu := newstudent(id, name)
   //追加到map中
   allStudent[id] = newStu
}
func delteStudent() {
   var deleteID int
   fmt.Println("请输入要删除的学号：")
   fmt.Scanln(&deleteID)
   delete(allStudent, deleteID)
}
func main() {
   allStudent = make(map[int]*student, 50)
   for {
      fmt.Println("欢迎来到学生管理系统")
      fmt.Println("请输入您的操作：")
      fmt.Println("1、查看所有学生信息")
      fmt.Println("2、添加学生")
      fmt.Println("3、删除学生信息")
      fmt.Println("4、退出学生信息管理系统")
      var choice int
      fmt.Scanln(&choice)
      switch choice {
      case 1:
         showStudent()
      case 2:
         addStudent()
      case 3:
         delteStudent()
      case 4:
         fmt.Println("再见！")
         os.Exit(1)
      default:
         fmt.Println("无效输入！")
      }
   }
}
```

### 结构体版

![image-20230306112115568](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230306112115568.png)

![image-20230306112312364](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230306112312364.png)

![image-20230306112639494](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230306112639494.png)

![image-20230306112535745](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230306112535745.png)  

![image-20230306112509409](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230306112509409.png)

![image-20230306112447475](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230306112447475.png)

### 区别：

函数版把所有的数据放到一个全局变量，所有的函数都操作那一个全局变量

结构体版把管理系统作为一个物件，给这个物件赋予数据和动作，定义了一个结构体管理者的数据还有一些方法

## 接口：是一种类型

```go
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
```

接口是一种特殊的类型，它规定了变量有哪些方法。

**场景**：不管传进来什么类型，我只关心它可以调用什么方法。

```go
type dog struct {
   name string
}
type cat struct {
   name string
}
// 方法
func (d dog) speak() {
   fmt.Printf("%v会汪汪汪~\n", d.name)
}
func (c cat) speak() {
   fmt.Println(("喵喵喵~"))
}
type speaker interface { //接口是一种类型
   speak() //接收到了什么方法
}
func da(x speaker) { //定义了一个名为da的函数，传入了一个变量，变量类型为接口类型
   x.speak() //这个接口类型的变量做了什么方法
}
func main() {
   var d1 dog
   d1.name = "大黄"
   var c1 cat
   da(d1) //大黄会汪汪汪~
   da(c1) //喵喵喵~
}
```

**接口的使用**：

1、定义结构体

2、定义结构体可以使用的方法

3、**定义一个接口**，接收到上述方法（接口是一种类型）

4、定义一个函数，传入接口类型，**调用接口的方法**

总结：结构体可以看作接口类型去调用这个方法。

```go
type baoshijie struct {
   brand string
}
type falali struct {
   brand string
}
func (b baoshijie) pao() {
   fmt.Printf("%s的速度是700迈\n", b.brand)
}
func (f falali) pao() {
   fmt.Printf("%s的速度是7000迈\n", f.brand)
}
type paoche interface { //定义接口     把执行这个方法的东西放到一起，定义为一个大类
   pao()      		   //可以执行的方法
}
func drive(p paoche) { //执行方法 的函数
   p.pao()
}
func main() {
   b1 := baoshijie{
      brand: "保时捷",
   }
   f1 := falali{
      brand: "法拉利",
   }
   b1.pao()//保时捷的速度是700迈
   f1.pao()//法拉利的速度是7000迈
}
```

### 接口的实现

 一个变量如果实现了接口中规定的所有方法，那么这个变量就实现了这个接口，可以成为这个接口类型的变量。

![image-20230306210501889](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230306210501889.png)

![image-20230306210629310](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230306210629310.png)

![image-20230306210645657](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230306210645657.png)

```
// 接口的实现
type catt struct {
   name string
   feet int8
}

func (c catt) move() {
   fmt.Println("猫猫出击")
}
func (c catt) eat(food string) {
   fmt.Printf("%s爱吃吃%s...\n", c.name, food)
}

type animal interface {
   move()
   eat(fo string)
}

// ****
func hunr(a animal) {
   a.move()
   a.eat(string("猫粮"))
}
func main() {
   var aa animal
   aa = catt{
      name: "米粒",
      feet: 8,
   }
   aa.move()     //猫猫出击
   aa.eat("鱼罐头") //米粒爱吃吃鱼罐头...

   //***
   hunr(catt{
      name: "花花",
   }) //猫猫出击
   //花花爱吃吃猫粮...
}
```
