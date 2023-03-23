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
