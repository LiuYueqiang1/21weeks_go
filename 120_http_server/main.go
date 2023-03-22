package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//str := "hello world!"
	str, err := ioutil.ReadFile("F:\\goland\\go_project\\21weeks\\21weeks_go\\81_rizhipackge_make\\consloe\\consloe.go")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("无内容,err,%v:\n", err)))
	}
	w.Write(str)
}
func main() {
	http.HandleFunc("/posts/go/test/", f1)     //目录
	http.ListenAndServe("127.0.0.1:9090", nil) //书
}
