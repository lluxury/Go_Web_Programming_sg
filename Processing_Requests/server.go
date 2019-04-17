package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request)  {
	r.ParseMultipartForm(1024)
	// fmt.Fprintln(w,r.MultipartForm)
	fmt.Fprintln(w,r.FormValue("hello"))
	fmt.Fprintln(w,r.Form)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}