package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func process(w http.ResponseWriter, r *http.Request)  {
	r.ParseMultipartForm(1024)

	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
           fmt.Fprintln(w, string(data))
		}
	}
}

// 执行 ParseMultipartForm ,从MultipartForm的File字段取出文件头 fileHeader
// 通过调用文件头的Open方法打开文件,读取到字节数组,并打印

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
