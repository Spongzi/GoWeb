package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func request(w http.ResponseWriter, r *http.Request) {
	// 这些信息是输出到服务器端的打印信息
	fmt.Println("Request 解析")
	// HTTP方法
	fmt.Println("method", r.Method)
	// RequestURI 是被客户端发送到服务器端的请求行中未修改的请求URI
	fmt.Println("RequestURI: ", r.RequestURI)
	// URL 类型，下方分别列出URL成员
	fmt.Println("URL_PATH", r.URL.Path)
	fmt.Println("URL_RawQuery", r.URL.RawQuery)
	fmt.Println("URL_Fragment", r.URL.Fragment)
	// 协议版本
	fmt.Println("proto", r.Proto)
	fmt.Println("protomajor", r.ProtoMajor)
	fmt.Println("protominor", r.ProtoMinor)
	// HTTP请求头
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Println("header key:" + k + " value:" + vv)
		}
	}
	// 判断是否为multipart
	isMultipart := false
	for _, v := range r.Header["Content-Type"] {
		if strings.Index(v, "multipart/form-data") != -1 {
			isMultipart = true
		}
	}
	// 解析From 表单
	if isMultipart == true {
		r.ParseMultipartForm(128)
		fmt.Println("解析方式：ParseMultipart")
	} else {
		r.ParseForm()
		fmt.Println("解析方式:Parse")
	}
	// HTTP Body 内容长度
	fmt.Println("ContentLength", r.ContentLength)
	// 是否在回复请求后关闭连接
	fmt.Println("Close", r.Close)
	// HOST
	fmt.Println("host", r.Host)
	// 该请求来源地址
	fmt.Println("RemoteAddr", r.RemoteAddr)
	fmt.Fprintf(w, "Hello, let's go!") // 这个是输出到客户端的
}

func main() {
	http.HandleFunc("/hello", request)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("ListenAndServe", err)
	}
}
