package main

// handler 使用
import "net/http"

// Refer 先创建结构体，只有带上指定的refer参数，该请求才能调用成功，否则返回403
type Refer struct {
	handler http.Handler
	refer   string
}

// 需要将这个Refer 实例化传给ListAndServer()函数，因此必须要实现ServeHTTP()方法，可以直接编写用来实现中间件的逻辑
func (this *Refer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Referer() == this.refer {
		this.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

// 取出当前请求头中的refer信息，如果与约定的不同，则拦截请求
func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is handler"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	referer := &Refer{
		handler: http.HandlerFunc(myHandler),
		refer:   "www.shiron.com",
	}
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", referer)
}
