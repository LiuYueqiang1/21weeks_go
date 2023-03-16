package main

import (
	"fmt"
	"time"
)

func timetest() {
	nowTime := time.Now()
	fmt.Println(nowTime)
	fmt.Println(nowTime.Year())
	fmt.Println(nowTime.Month())
	fmt.Println(nowTime.Day())
	fmt.Println(nowTime.Hour())
	fmt.Println(nowTime.Minute())
	fmt.Println(nowTime.Second()) //秒
	time1 := nowTime.Unix()       //时间戳
	time2 := nowTime.UnixNano()   //纳秒时间戳
	fmt.Println(time1, time2)
	a1 := time.Unix(1678886401, 0)
	fmt.Println(a1) //根据时间戳查看时间
}
func timezoneDemo() {
	// 中国没有夏令时，使用一个固定的8小时的UTC时差。
	// 对于很多其他国家需要考虑夏令时。
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	// FixedZone 返回始终使用给定区域名称和偏移量(UTC 以东秒)的 Location。
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	fmt.Println(beijing)
}

// 时间戳
func timestampDemo() {
	now := time.Now()
	fmt.Println(now)                  //2023-03-09 20:12:21.3150337 +0800 CST m=+0.007772701
	tafter_hour := now.Add(time.Hour) //加一个小时
	fmt.Println(tafter_hour)          //2023-03-09 21:12:21.3150337 +0800 CST m=+3600.007772701
	timeunix := now.Unix()            //微秒时间戳
	timemilli := now.UnixMilli()      //毫秒时间戳
	timemicro := now.UnixMicro()      //微秒时间戳
	timenano := now.UnixNano()        //纳秒时间戳
	fmt.Println(timeunix, timemilli, timemicro, timenano)
}

// timestamp2Time  将时间戳转化为对象
func timestamp2Time() {
	//获取北京时间所在的东八区对象
	//secondsEastofUtc := int((8 * time.Hour).Seconds())
}

// 定时器
// 使用time.Tick(时间间隔)来设置定时器，定时器的本质上是一个通道（channel）
func tickDemo() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Printf("%T\n", ticker) //<-chan time.Time
		fmt.Println(i)             //每秒都会执行的任务
	}
}

func fomatdemo() {
	nowTime := time.Now()
	ret1 := nowTime.Format("2006-01-02 15:04:05.000 Mon Jan")
	fmt.Println(ret1)
	ret2 := nowTime.Format("2006-01-02 03:04:05.000 PM Mon Jan")
	fmt.Println(ret2)
	fmt.Println(nowTime.Format("2006-01-02"))
	fmt.Println(nowTime.Format("15:04:05"))
}
func parseDemo() {
	//在没有时区指示符的情况下，time.Prase 返回UTC时间
	timeObj, err := time.Parse("2006-01-02 15:04:05.000 Mon Jan", "2023-03-09 20:53:29.342 Thu Mar")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj) //2023-03-09 20:53:29.342 +0000 UTC
	// 在有时区指示符的情况下，time.Parse 返回对应时区的时间表示
	// RFC3339     = "2006-01-02T15:04:05Z07:00
	timeObj2, err2 := time.Parse(time.RFC3339, "2022-10-05T11:25:20+08:00")
	if err2 != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj2) //2023-03-09 20:53:29.342 +0000 UTC
}

// time.ParseInLocation函数需要在解析时额外指定时区信息。
func parseDemo2() {
	now := time.Now()
	fmt.Println(now)
	//加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2022/10/05 11:25:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}
func main() {
	//timetest()
	//timezoneDemo()
	//timestampDemo()
	//fomatdemo()
	//tickDemo()
	parseDemo2()
	//n := 5
	//fmt.Println("开始sleep")
	//time.Sleep(time.Duration(n) * time.Second)
	//fmt.Println("5秒钟过去了")
	//time.Sleep(2 * time.Second)
	//fmt.Println("又2秒钟过去了")
}
