package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 创建一个GET请求
func main() {
	resp, err := http.Get("https://www.bilibili.com")
	if err != nil {
		fmt.Println("err", err)
	}
	closer := resp.Body
	bytes, err := ioutil.ReadAll(closer)
	fmt.Println(string(bytes))
}
