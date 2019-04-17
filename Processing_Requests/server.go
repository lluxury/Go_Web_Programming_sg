package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request)  {
	// r.ParseMultipartForm(1024)
	// r.ParseForm
	// fmt.Fprintln(w,r.PostFormValue("hello"))
	// fmt.Fprintln(w,r.PostForm)

	fmt.Fprintln(w, "(1)", r.FormValue("hello"))
	fmt.Fprintln(w, "(2)", r.PostFormValue("hello"))
	fmt.Fprintln(w, "(3)", r.PostForm)
	fmt.Fprintln(w, "(4)", r.MultipartForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}