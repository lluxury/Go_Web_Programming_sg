package main

import (
	"net/http"
	"fmt"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello yann!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "World yann!")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// Handler: &handler,
	}
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	
	server.ListenAndServe()
}