package main

import (
	"fmt"
	"log"
	"net/http"
)

func hi(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "hi")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hi", hi)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
		return
	}
}
