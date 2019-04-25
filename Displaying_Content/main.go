package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmp1.html")
	// t.Execute(w, content)
	t.Execute(w, r.FormValue("comment"))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	// t.Execute(w, nil)
	t.Execute(w, template.HTML(r.FormValue("comment")))

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/form", form)

	server.ListenAndServe()
}
