package main

import "net/url"

func main() {
	path := "http://localhost:8080/article?id=1"
	p, _ := url.Parse(path)
	println(p.Host)
	println(p.User)
	println(p.RawQuery)
	println(p.RequestURI())
}
