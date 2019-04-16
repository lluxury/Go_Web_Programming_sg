package main

import (
	"net/http"
	"fmt"
)

type HelloHandler struct {}

func (h HelloHandler) ServeHTTP (w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"Hello")
}

func log(h http.Handler) http.Handler  {
	return http.HandlerFunc (func (w http.ResponseWriter, r *http.Request)  {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)		
	})
	
}

func protect(h http.Handler) http.Handler  {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
	// some code
	fmt.Printf("protect work - %T\n", h)
	h.ServeHTTP(w, r)		
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello := HelloHandler{}
	// http.HandleFunc("/hello", protect(log(hello)))
	http.Handle("/hello", protect(log(hello)))
	
	server.ListenAndServe()
}