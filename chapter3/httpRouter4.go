package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

//  静态文件服务
func main() {
	router := httprouter.New()

	// 访问静态文件
	router.ServeFiles("/static/*filepath", http.Dir("."))

	// 启动服务器
	log.Fatal(http.ListenAndServe(":8080", router))
}
