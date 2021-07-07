package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "apple apple apple")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "grape grape grape")
}

func main() {

	http.HandleFunc("/apple", d)
	http.HandleFunc("/grape", c)

	http.ListenAndServe(":8080", nil)
}
