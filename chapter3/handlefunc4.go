package main

import (
	"fmt"
	"log"
	"net/http"
)

// 同时使用处理器和处理函数

func hiHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi Go handleFunc")
}

type welcomeHandler struct {
	Name string
}

func (h welcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi %s", h.Name)
}

func main() {
	mux := http.NewServeMux()

	// 注册处理器函数
	mux.HandleFunc("/hi", hiHandler1)

	// 注册处理器
	mux.Handle("/welcome/goweb", welcomeHandler{
		Name: "Go handle",
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
