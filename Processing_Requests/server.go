package main

import (
	"encoding/base64"
	"net/http"
	"fmt"
	"time"
)

// func setCookie(w http.ResponseWriter, r *http.Request)  {
func setMessage(w http.ResponseWriter, r *http.Request)  {
	msg := []byte("Hello World!")

	c := http.Cookie{
		Name: "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
		// 编码,避免 空格和特殊字符
	}
	http.SetCookie(w, &c)
	// 传递给方法的是结构体指针,而不是本身 &c1 指针
}

func showMessage(w http.ResponseWriter,r *http.Request )  {
	c, err := r.Cookie("flash")
	if err != nil {
			if err == http.ErrNoCookie{
			fmt.Fprintln(w, "No message found")
		  	}
		} else {
			rc := http.Cookie{
				Name: "flash",
				MaxAge: -1,
				Expires: time.Unix(1,0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		// 获取值,过期,解码
		fmt.Fprintln(w, string(val))
	}
}
// 注意缩进格式,不然很难排查

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}
