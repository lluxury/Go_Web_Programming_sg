package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string  {
	layout := "2019-04-25"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {

	// funcMap := template.FuncMap("fdate": formatDate)
	funcMap := template.FuncMap{"fdate": formatDate}
	
	t :=template.New("tmp1.html").Funcs(funcMap)
	// t, _ := template.ParseFiles("t1.html", "t2.html")
	
	t, _ = t.ParseFiles("tmp1.html")
	// t.Execute(w, "Hello World!")
	
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
