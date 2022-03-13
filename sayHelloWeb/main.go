package main

import "net/http"

// 创建一个简单的Go HTTP 服务器
func sayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	http.ListenAndServe(":8080", nil)
}

/*
	内部逻辑分析：
	如果要创建一个Web服务器端，则要①调用http.HandleFunc() 函数； ②调用http.ListenAndServe()函数
	ListAndServe()函数有两个参数：当前监听的端口号和事件处理器handler
*/
