package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", lion)
	http.ListenAndServe(":8080", nil)
}

func lion(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!--image doesn't serve-->
	<img src="/lion.webp">
	`)
}
