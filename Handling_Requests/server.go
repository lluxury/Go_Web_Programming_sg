package main

import (
	"net/http"
	"fmt"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello World yann!")
}
func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
	// server.ListenAndServeTLS("cert.pem","key.pem")
}