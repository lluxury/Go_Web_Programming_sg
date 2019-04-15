package main

import (
	"net/http"
	"html/template"
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
    
    func index(w http.ResponseWriter,r *http.Request)  {
     // files := []string{"templates/layout.html",
     //              "templates/navbar.html",
     //              "templates/index.html",}
     // templates := template.Must(template.ParseFiles(files...))
     threads, err := data.threads();if err == nil{
         _, err := session(w, r)
         public_tmpl_files :=  []string{"templates/layout.html",
                                      "templates/public.navbar.html",
                                      "templates/index.html",}
         private_tmpl_files :=  []string{"templates/layout.html",
                                         "templates/private.navbar.html",
                                         "templates/index.html",}
         var template *template.Template
         if err !=nil {
             templates = template.Must(template.ParseFiles(private_tmpl_files...))
         } else {
             templates = template.Must(template.ParseFiles(public_tmpl_files...))
         }
         templates.ExeuteTemplate(w,"layout", threads)
     }
    }