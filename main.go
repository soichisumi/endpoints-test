package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func main() {
	http.Handle("/hello", String("Hello, World"))
	http.ListenAndServe("0.0.0.0:8080", nil)
}