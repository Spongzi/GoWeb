package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./hello.html")
	if err != nil {
		fmt.Println("create template failed, err: ", err)
		return
	}
	err = tmpl.Execute(w, UserInfo{
		Name:   "李四",
		Gender: "男",
		Age:    28,
	})
	if err != nil {
		fmt.Println("template execute failed, err: ", err)
		return
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	panic(http.ListenAndServe(":8080", nil))
}
