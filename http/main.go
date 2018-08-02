package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w io.Writer, name string) {
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "hello,%s", name)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hello(w, "world")
}

func main() {
	http.ListenAndServe(
		":54321",
		http.HandlerFunc(helloHandler),
	)
}
