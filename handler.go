package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// IndexHandler is handler function for "/"
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := ioutil.ReadFile("index.html")
	w.Write(page)
}

// HelloHandler is handler function for "/hello"
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}
