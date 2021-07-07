package main

import (
	"io"
	"net/http"
)

type apple int

func (m apple) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy doggy doggy")
	case "/cat":
		io.WriteString(w, "kitty kitty kitty")
	}
}

func main() {
	var d apple
	http.ListenAndServe(":8080", d)
}
