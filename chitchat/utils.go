package main

import (
	"net/http"
	"yann/Go_Web_Programming_sg/chitchat/data"
	"errors"
	"html/template"
	"fmt"
	
)
func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error)  {
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
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		// files = append(falses, fmt.Sprintf("templates/%s.html", file))
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
	
}

