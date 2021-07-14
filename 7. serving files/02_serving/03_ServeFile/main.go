package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", lion)
	http.HandleFunc("/lion.webp", lionPic)
	http.ListenAndServe(":8080", nil)
}

func lion(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="lion.webp">	`)
}

func lionPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "lion.webp")
}
