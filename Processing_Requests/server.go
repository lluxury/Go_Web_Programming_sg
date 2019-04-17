package main

import (
	"fmt"
	"net/http"
)

// func body(w http.ResponseWriter, r *http.Request)  {
func process(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	// fmt.Fprintln(w,r.Form)
	// fmt.Fprintln(w,r.Form["hello"])
	// fmt.Fprintln(w,r.PostForm)
	fmt.Fprintln(w,r.Form)
}
// 使用 ParseForm对请求进行语法分析,再访问Form字段,获取表单

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// http.HandleFunc("/body", body)
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}