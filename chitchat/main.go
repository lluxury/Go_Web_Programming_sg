package main

import (
	"net/http"
	"html/template"
	"yann/Go_Web_Programming_sg/chitchat/data"

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

func index(w http.ResponseWriter, r *http.Request) {
  files := []string{"templates/layout.html",
                    "templates/navbar.html",
                    "templates/index.html",}
  templates := template.Must(template.ParseFiles(files...))
  threads, err := data.Threads(); if err == nil {
    templates.ExecuteTemplate(w, "layout", threads)
  }
}