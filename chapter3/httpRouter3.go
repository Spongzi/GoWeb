package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// Handler包可以处理不同的二级域名，它现根据域名湖区对应的Handler路由，然后调用处理(分发机制)

type HostMap map[string]http.Handler

func (hs HostMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 先根据域名获取对应的Handler路由，然后调用处理(分发机制)
	if handler := hs[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

func main() {
	userRouter := httprouter.New()

	// 根路径
	userRouter.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		if _, err := writer.Write([]byte("sub1")); err != nil {
			return
		}
	})

	dataRouter := httprouter.New()

	dataRouter.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		if _, err := writer.Write([]byte("sub2")); err != nil {
			return
		}
	})

	// 分别处理不同的二级域名
	hs := make(HostMap)
	hs["sub1.localhost:8080"] = userRouter
	hs["sub2.localhost:8080"] = dataRouter

	log.Fatal(http.ListenAndServe(":8080", hs))
}
