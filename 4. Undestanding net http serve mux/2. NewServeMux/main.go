package main

import (
	"io"
	"net/http"
)

type apple int

func (d apple) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "apple apple apple")
}

type pineapple int

func (c pineapple) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "pineapple pineapple pineapple")
}

func main() {
	var d apple
	var c pineapple

	mux := http.NewServeMux()
	mux.Handle("/apple/", d)
	mux.Handle("/pineapple", c)

	http.ListenAndServe(":8080", mux)
}
