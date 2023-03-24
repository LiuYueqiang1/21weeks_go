package main

import "fmt"

func main() {
	var si = make([]int, 0, 10)
	si = []int{1, 3, 5, 6, 8, 5, 4, 2}
	si = append(si[:2], si[3:]...)
	fmt.Println(si) //[1 3 6 8 5 4 2]
	fmt.Println(si[1])

}
