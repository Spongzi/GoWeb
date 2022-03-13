package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "欢迎来到Go Web 首页！ 处理器为:indexHandler!")
}

func hiHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "欢迎来到Go Web 首页！ 处理器为:hiHandler!")
}

func webHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "欢迎来到Go Web 首页！ 处理器为:webHandler!")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/hi", hiHandler)
	mux.HandleFunc("/hi/web", webHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
		return
	}
}
