package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Index"))
}

func main() {
	r := httprouter.New()
	r.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", r))
}
