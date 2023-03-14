package main

import (
	"fmt"
	"math/rand"
	"time"
	//	"time"
)

func t1() {
	for i := 0; i < 5; i++ {
		a := rand.Int()
		fmt.Println(a)
	}
}
func t2() {
	rand.Seed(time.Now().UnixNano()) //似乎已启用，不需要随机数种子了
	for i := 0; i < 5; i++ {
		r1 := rand.Int() //int64
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}
func t3() {
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
}
func main() {

	t3()

}
