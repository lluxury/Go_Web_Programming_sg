package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request)  {
	// h := r.Header
	// h := r.Header["Accept-Encoding"]
	h := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w,h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}