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
	log.Print("Starting Server on", port)
	err := http.ListenAndServe(port,mux)
	log.Fatal(err)
}
