package main

import (
	"log"
	"net/http"
)

const (
	port = ":7000"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting Server on", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
