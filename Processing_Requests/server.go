package main

import (
	"fmt"
	"net/http"
	// "io/ioutil"
)

func writeExample(w http.ResponseWriter, r *http.Request)  {
	str := `<html>
<head><title> Go Web Programming</title></head>
<body><h1>hello World yann</h1></body>
</html>	`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter , r *http.Request)  {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	server.ListenAndServe()
}
