package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Post struct {
	User string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request)  {
	str := `<html>
<head><title> Go Web Programming</title></head>
<body><h1>hello World yann</h1></body>
</html>	`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter , r *http.Request)  {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
	
}

func jsonExample(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User: "yann",
		// Threads: []string["first", "second", "third"],
		Threads: []string{"first", "second", "third"},
	}
	// 获取地址,赋值元素
	
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json",jsonExample)
	server.ListenAndServe()
}
