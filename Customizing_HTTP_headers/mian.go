package main

import (
	"log"
	"net/http"
)

const (
	port = ":7000"
)

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	//User r.Method to check whether the request is using POST or not.
	if r.Method != "POST" {
		//Use the w.Header.Set() method to add an 'Allow: POST' header to the
		//response header map. The first parameter is header name,and
		//the second parameter is the header value.
		w.Header().Set("Allow", "POST")
		//Use the http.Error() function to send a 405 status code and "Method Not
		//Allowed" string as the response body.
		http.Error(w,"Method Not Allowed",http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting........", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
