package main

import (
	"fmt"
	"time"
)

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
