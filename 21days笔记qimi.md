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

new()

```go
func p3() {
   var a = new(int)
   fmt.Println(a)  //0xc00009e058
   fmt.Println(*a) //0
   //Go里面的指针只能读不能修改，不能修改指针变量指向的地址
   *a = 100
   fmt.Println(a)  //0xc00009e058
   fmt.Println(*a) //100
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

## 闭包

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
