package main
import (
	"net/http"
	"yann/Go_Web_Programming_sg/chitchat/data"
)

func authenticate(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	user, _ := date.UserByEmail(r.PostFormValue("email"))
	// 给定地址,获取user结构,用户要先存在
	if user.Password == data.Encrypt(r.PostFormValue("password")){
		// 加密给定字符串
		session := user.CreateSession()
		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w,r,"/",302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
	
}