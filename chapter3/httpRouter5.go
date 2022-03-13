package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// HttpRouter包 允许使用这设置PanicHandler, 以处理再HTTP请求中发生的panic异常。

func Index1(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	panic("error")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index1)
	// 捕获异常
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, i interface{}) {
		writer.WriteHeader(http.StatusInternalServerError)
		if _, err := fmt.Fprintf(writer, "error: %s", i); err != nil {
			return
		}
	}
	log.Fatal(http.ListenAndServe(":8080", router))
}
