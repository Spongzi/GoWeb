package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	// 1. 解析模板
	t, err := template.ParseFiles("./template_example.tmpl")
	if err != nil {
		fmt.Println("template parseFiles failed, err:", err)
		return
	}
	// 2. 渲染模板
	name := "我爱Go语言"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("template execute failed, err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", helloHandleFunc)
	panic(http.ListenAndServe(":8086", nil))
}
