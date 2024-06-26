package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From Snippetbox"))
}

func main() {
	//use the http.NewServerMux() function to initialize a new servermux then
	//register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	//Use the http.ListenAndServe() function to start a new web server. We pass in
	//two parameters: the TCP network address to listen on (in this case ":7000")
	//and the servemux we just created. If http.ListenAndServe() returns an error
	//we use the log.Fatal() function to log the error message and exit . Note
	//that any error returned by http.ListenAndServe() is always non-nil
	log.Print("Starting serve on 7000")
	err := http.ListenAndServe(":7000", mux)
	log.Fatal(err)
}
