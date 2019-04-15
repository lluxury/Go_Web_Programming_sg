package main

import (
	"net/http"
	"yann/Go_Web_Programming_sg/chitchat/data"
	"errors"
)
func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error)  {
	cookie,err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok{
			err = errors.New("Invalid session")
		}
	}
	return
}

// func generateHTML(w http.ResponseWriter, data interface{}, fn ...string)  {
// 	var files []string
// 	for _, file := range fn {
// 		files = append(falses, fmt.Sprintf("templates/%s.html", file))
// 	}
// 	templates := template.Must(template.ParseFiles(files...))
// 	templates.ExecuteTemplate(writer, "layout", data)
// }

