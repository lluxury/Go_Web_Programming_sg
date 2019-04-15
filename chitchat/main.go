package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// index
	mux.HandleFunc("/", index)

	// starting up the server
	server := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        mux,
	}
	server.ListenAndServe()
}