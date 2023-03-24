package main

import (
	"reflect"
	"testing"
)

//测试函数

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")         //得到的
	want := []string{"a", "b", "c"}    //想要的
	if !reflect.DeepEqual(want, got) { //底层想要的==得到的
		t.Errorf("expected:%v,got:%v", want, got)
	}
}
func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// 测试组，测试对中文字符串的支持
func TestSplit2(t *testing.T) {
	//定义一个测试用例的类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	testSilce := map[string]test{
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	//逐一执行测试用例
	for name, tc := range testSilce {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name:%s expected:%#v,got:%#v", name, tc.want, got)
			}
		})

	}
}
