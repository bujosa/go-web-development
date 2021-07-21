package main

import (
	"io"
	"net/http"
)

type apple int

func (d apple) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "apple apple apple")
}

type grape int

func (c grape) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "grape grape grape")
}

func main() {
	var d apple
	var c grape

	http.Handle("/apple", d)
	http.Handle("/grape", c)

	http.ListenAndServe(":8080", nil)
}
