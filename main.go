package main

import (
	"fmt"
	"net/http"

	"github.com/jaguan/webserver/handler"
)

func main() {
	// 	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 		w.Write([]byte("Hello World"))
	// 	})

	// staticHandler := handler.GetStaticHandler()
	// http.Handle("/", &staticHandler)
	http.Handle("/", new(handler.StaticHandler))

	fmt.Println("Server Start")
	http.ListenAndServe(":5000", nil)
	fmt.Println("Server End")
}

func init() {
	fmt.Println("main Start")
}
