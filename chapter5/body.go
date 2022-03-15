package main

import (
	"fmt"
	"net/http"
)

func getBody(w http.ResponseWriter, r *http.Request) {
	// 获取请求报文的内容长度
	contentLength := r.ContentLength
	// 新建一个字节切片， 长度与请求报文的内容和长度相同
	body := make([]byte, contentLength)
	// 读取r的一个请求体，并将具体内容写入Body当中
	r.Body.Read(body)
	// 将获取的参数内容写入相应的报文中
	fmt.Fprintf(w, string(body))
}

func main() {
	http.HandleFunc("/getbody", getBody)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}
