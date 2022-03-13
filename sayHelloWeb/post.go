package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url := "https://www.shirdon.com/comment/add"
	body := "{\"userId\":1,\"articleId\":1,\"comment\":\"这是一条评论\"}"
	response, err := http.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println("err", err)
	}
	b, err := ioutil.ReadAll(response.Body)
	file, err := os.OpenFile("blog.html", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	_, err = file.Write(b)
	if err != nil {
		return
	}
	fmt.Println(string(b))
}
