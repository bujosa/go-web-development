package main

import (
	"io"
	"net/http"
)

type apple int

func (m apple) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/apple":
		io.WriteString(w, "apple apple apple")
	case "/pineapple":
		io.WriteString(w, "pineapple pineapple pineapple")
	}
}

func main() {
	var d apple
	http.ListenAndServe(":8080", d)
}
