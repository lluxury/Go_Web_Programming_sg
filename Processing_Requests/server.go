package main

import (
	"net/http"
	"fmt"
)

func setCookie(w http.ResponseWriter, r *http.Request)  {
	c1 := http.Cookie{
		Name: "first_cookie",
		Value: "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name: "second_cookie",
		Value: "yann's test",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
	// 传递给方法的是结构体指针,而不是本身 &c1 指针
}

func getCookie(w http.ResponseWriter,r *http.Request )  {
	h := r.Header["Cookie"]
	fmt.Fprintln(w,h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie",setCookie)
	http.HandleFunc("/get_cookie",getCookie)
	server.ListenAndServe()
}
