package main

import (
	"net/http"
	// "html/template"
	"yann/Go_Web_Programming_sg/chitchat/data"

)

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// index
	mux.HandleFunc("/", index)
// 发送到根URL的请求重定向到服务器  参数 URL, 处理器名,没有显式写index的参数

	// starting up the server
	server := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        mux,
	}
	server.ListenAndServe()
}

    // index
// func index(write http.ResponseWriter, request *http.Request)  {
    func index(writer http.ResponseWriter, request *http.Request) {
        // threads, err := data.threads(); if err != nil {
		threads, err := data.Threads()
		if err != nil {
			error_message(writer,request,"Cannot get threads")
		} else {
            _, err := session(writer, request)
            if err !=nil{
                generateHTML(writer, threads, "layout", "public.navbar", "index")
        } else {
                generateHTML(writer, threads, "layout", "private.navbar", "index")
        }
    }
}