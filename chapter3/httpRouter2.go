package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	// 创建一个router
	router := httprouter.New()

	// 书写默认的Get方法
	router.GET("/default", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		if _, err := writer.Write([]byte("default get")); err != nil {
			return
		}
	})

	// 书写默认的POST请求
	router.POST("/default", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		if _, err := writer.Write([]byte("default post")); err != nil {
			return
		}
	})

	// 精准匹配
	router.GET("/user/:name", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		if _, err := writer.Write([]byte("user name: " + params.ByName("name"))); err != nil {
			return
		}
	})

	// 匹配所有
	router.GET("/user/*name", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		if _, err := writer.Write([]byte("user name: " + params.ByName("name"))); err != nil {
			return
		}
	})

	// 启动服务器
	panic(http.ListenAndServe(":8080", router))
}
