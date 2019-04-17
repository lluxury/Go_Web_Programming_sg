package main

import (
	"fmt"
	"net/http"
)

// func headers(w http.ResponseWriter, r *http.Request)  {
func body(w http.ResponseWriter, r *http.Request)  {
	// len := ContentLength
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w,string(body))
}
// 获取主体数据长度,创建字节数组,调用Read方法将数据读入数组

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// http.HandleFunc("/headers", headers)
	http.HandleFunc("/body", body)
	server.ListenAndServe()
}