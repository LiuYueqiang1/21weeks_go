package main

import (
	"fmt"
)

// 复习接口基础1
type Wechat struct {
	name string
}

// 方法
func (w Wechat) zhifu() {
	fmt.Println("微信支付")
}

type zhifubao struct {
	name string
}

func (z zhifubao) zhifu() {
	fmt.Println("支付宝支付")
}

type fuqian interface {
	zhifu()
}

// *****//
func futype(f fuqian) {
	f.zhifu()
}

func main() {
	w1 := Wechat{
		name: "微信",
	}
	z1 := zhifubao{
		name: "支付宝",
	}
	var f1 fuqian
	f1 = w1
	f1.zhifu() //微信支付
	f1 = z1
	f1.zhifu() //支付宝支付

	//*****
	futype(w1) //微信支付
	futype(z1) //支付宝支付
}
