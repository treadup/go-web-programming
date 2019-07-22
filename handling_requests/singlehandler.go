package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (*HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := HelloHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: &handler,
	}

	server.ListenAndServe()
}
