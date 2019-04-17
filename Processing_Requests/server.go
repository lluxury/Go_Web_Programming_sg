package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request)  {
	// r.ParseForm()
	r.ParseMultipartForm(1024)
	// fmt.Fprintln(w,r.Form)
	fmt.Fprintln(w,r.MultipartForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}