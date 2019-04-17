package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request)  {
	// r.ParseMultipartForm(1024)
	// r.ParseForm
	fmt.Fprintln(w,r.PostFormValue("hello"))
	fmt.Fprintln(w,r.PostForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}