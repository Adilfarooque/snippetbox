package main

import (
	"log"
	"net/http"
)

const (
	port = ":7000"
)

func home(w http.ResponseWriter, r *http.Request) {
	//Check if the current request URL path exactly matches "/" . If it donsn't, use
	//the http.NotFound() function to send a 404 to the client.
	//Importantly , we then  return from the handler. If we don't return the handler
	//would keep executing and also write the "Hello from SnippetBox" message.
	if r.URL.Path != "/"{
		http.NotFound(w,r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetview handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add a snippetCreate handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	//Register the two new handler functions and corresponding URL patters with
	//the servermux , in exactly the same way that we did before
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view/", snippetView)
	mux.HandleFunc("/snippet/create/", snippetCreate)

	log.Print("Starting server on", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
