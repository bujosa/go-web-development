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

	http.Handle("/apple", http.HandlerFunc(d))
	http.Handle("/grape", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}

// this is similar to this:
// https://play.golang.org/p/X2dlgVSIrd
// ---and this---
// https://play.golang.org/p/YaUYR63b7L
