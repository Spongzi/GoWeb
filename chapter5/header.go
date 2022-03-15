package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Greeting struct {
	Message string `json:"message"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	greeting := Greeting{
		Message: "欢迎一起学习《GoWeb编程实战派》",
	}
	marshal, _ := json.Marshal(greeting)
	// 通过set方法设置Content-Type为application/json类型
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshal)
}

func main() {
	http.HandleFunc("/hello", Hello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
		return
	}
}
