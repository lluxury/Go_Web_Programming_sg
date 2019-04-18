package main

import (
	// "fmt"
	"net/http"
	// "io/ioutil"
)

// func process(w http.ResponseWriter, r *http.Request)  {
func writeExample(w http.ResponseWriter, r *http.Request)  {
	str := `<html>
<head><title> Go Web Programming</title></head>
<body><h1>hello World yann</h1></body>
</html>	`
	w.Write([]byte(str))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// http.HandleFunc("/process", process)
	http.HandleFunc("/write", writeExample)
	server.ListenAndServe()
}
