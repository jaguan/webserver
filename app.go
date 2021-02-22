package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// App is ...
type App struct {
	Router *mux.Router
}

// Initialize is ...
func (a *App) Initialize() {
	fmt.Println("[APP] Initialize")

	a.Router = mux.NewRouter().StrictSlash(true)

	a.Router.HandleFunc("/", IndexHandler)
	a.Router.HandleFunc("/hello", HelloHandler)
}

// Run is ...
func (a *App) Run(port int) {
	fmt.Println("[APP] Run")

	var addr string = ":" + strconv.Itoa(port)

	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
