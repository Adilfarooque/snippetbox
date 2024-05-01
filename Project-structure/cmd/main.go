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
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet/view", SnippetView)
	mux.HandleFunc("/snippet/create", SnippetCreate)

	log.Print("Starting Server on", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
