package main

import (
	"net/http"
	"fmt"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"Hello")
}

func log(h http.HandlerFunc) http.HandlerFunc  {
			// 接收 HandleFunc 类型参数
	return func(w http.ResponseWriter, r *http.Request)  {
			// log 返回 HandleFunc 类型函数
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
			// 返回函数获取传入函数名字,再调用函数(h参数)
		fmt.Println("Handler function called -" + name)
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", log(hello))
	
	server.ListenAndServe()
}