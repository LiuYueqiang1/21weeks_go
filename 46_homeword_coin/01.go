package main

import (
	"fmt"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	//dispatchCoin()
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}
func dispatchCoin() (left int) {
	//1、拿到所有的用户，用for
	//2、每个用户中查看是否有eiou，EIOU这些字母
	//3、如果有的话，按照每个字母分金币 用value,ok=map[v]  ，map[name]coin 中对应的coin +
	//4、for每个map，输出对应的名字和金币 n,v:=for range map
	//5、获取所有的金币数量num，剩余的50-num

	for _, v := range users {
		if _, ok := distribution[v]; !ok {
			//3、将遍历得到的切片放到一个map中
			distribution[v] = 0
		}
	}
	fmt.Println(distribution)
	for n := range distribution { //拿到的n是每个人的名字
		//	fmt.Println(n)
		//fmt.Println("a ", distribution[n])
		for _, v2 := range n { //将每个人的名字里的字母全部取出来
			v3 := string(v2) //将其转换为string类型
			//fmt.Println("V3:", v3)
			switch {
			case v3 == "e":
				distribution[n] = distribution[n] + 1
			case v3 == "E":
				distribution[n] = distribution[n] + 1
			case v3 == "i":
				distribution[n] = distribution[n] + 2
			case v3 == "I":
				distribution[n] = distribution[n] + 2
			case v3 == "o":
				distribution[n] = distribution[n] + 3
			case v3 == "O":
				distribution[n] = distribution[n] + 3
			case v3 == "u":
				distribution[n] = distribution[n] + 4
			case v3 == "U":
				distribution[n] = distribution[n] + 4
			}
		}
	}
	sum := 0
	for _, v := range distribution {
		sum = sum + v
	}
	left = 50 - sum
	//_, ok1 := distribution["Aaron"]
	//if ok1 {
	//	distribution[n] = distribution[n] + 1
	//}
	//_, ok2 := distribution["E"]
	//if ok2 {
	//	distribution[n]++
	//}
	//_, ok3 := distribution["i"]
	//if ok3 {
	//	distribution[n] = distribution[n] + 2
	//}
	//_, ok4 := distribution["I"]
	//if ok4 {
	//	distribution[n] = distribution[n] + 2
	//}
	//_, ok5 := distribution["o"]
	//if ok5 {
	//	distribution[n] = distribution[n] + 3
	//}
	//_, ok6 := distribution["O"]
	//if ok6 {
	//	distribution[n] = distribution[n] + 3
	//}
	//_, ok7 := distribution["u"]
	//if ok7 {
	//	distribution[n] = distribution[n] + 4
	//}
	//_, ok8 := distribution["U"]
	//if ok8 {
	//	distribution[n] = distribution[n] + 4
	//}

	fmt.Println(distribution)
	return
	//1.依次拿到每个人的名字
	//2.拿到一个人名根据分金币的规则取分金币
	//2.1每个人分的金币数应该保存道distribution中
	//2.2记录下剩余的金币数
	//3.整个第2步执行完就能得到最终每个人分的金币数和剩余金币数
}
