package main

import (
	"html/template"
	"net/http"
	"math/rand"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmp1.html")
	rand.Seed(time.Now().Unix())
	daysOfWeek :=[]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	// t.Execute(w, rand.Intn(10) > 5)
	t.Execute(w, daysOfWeek)
}

func main() {
	server := http.Server{
		Addr : "127.0.0.1:8080",
	}
	http.HandleFunc("/process",process)
	server.ListenAndServe()
}
