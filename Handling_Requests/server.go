package main

import (
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"Hello")
}

func world(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"World")
}

// type HelloHandler struct{}

// func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
// 	fmt.Fprintf(w, "Hello yann!")
// }

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// http.Handle("/hello", &hello)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	
	server.ListenAndServe()
}