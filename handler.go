package main

import (
	"fmt"
	"net/http"
)

// IndexHandler is handler function for "/"
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index")
}

// HelloHandler is handler function for "/hello"
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
