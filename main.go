package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func view(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow","POST")
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed!"))
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/snippetcreated", snippetCreate)
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", view)
	mux.HandleFunc("/snippet/create", create)

	log.Print("Start server on :4000")
	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatal(err)
	}

}
