package main

import (
	"fmt" //New import
	"net/http"
	"strconv" //New import
)

const (
	port = ":7000"
)

func snippetView(w http.ResponseWriter, r *http.Request) {
	//Extract the value of the id parameter from the query string and try to
	//convert it to an integer using the strconv.Atoi() function.If it can't
	//be converted to an interger, or the value is less than 1 , we return a 404 page
	//not found response
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//Use the fmt.Fprintf() function to interpolate the id value with our response 
	//and write it to the http.ResponseWriter
	fmt.Fprintf(w, "Display the specific with ID %d", id)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/snippet/view", snippetView)
	http.ListenAndServe(port, mux)
}
