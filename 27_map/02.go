package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())        //初始化随机数种子
	var scoremap = make(map[string]int, 50) //制作分数map
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i) //每次循环都输出一个i
		value := rand.Intn(50)           //每次循环都生成一个0~50之间的随机数
		scoremap[key] = value
	}
	fmt.Println(scoremap)
	//取出所有的key存入切片
	keys := make([]string, 0, 100)
	for key := range scoremap {
		fmt.Println(key)
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的顺序遍历切片，输出切片名字和对应的map中的值
	for _, v := range keys {
		fmt.Println(v, scoremap[v])
	}
}
