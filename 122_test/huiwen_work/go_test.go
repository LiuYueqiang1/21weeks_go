package main

import "testing"

func TestPalindrome(t *testing.T) {
	err := Palindrome("一二三二一")
	if err != nil {
		t.Error(err)
	}
}

// 并行测试，创建多个goroutine，并将b.N分配给这些goroutine执行
func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Palindrome("一二三二一")
	}
}
