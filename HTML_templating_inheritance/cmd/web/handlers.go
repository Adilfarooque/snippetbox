package main

import (
	"fmt"
	"html/template" //New import
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//Use the template.ParseFiles() function to read the template file into a
	//template set. If there's an error, we long the detailed error message and use
	//the http.Error() function to send a generic 500 Internal Server Error
	//response to the user.
	tpl, err := template.ParseFiles("C:/Users/Adilf/OneDrive/Desktop/Code World/code/snippetbox./HTML_templating_inheritance/internal/ui/html/pages/home.tmpl")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error :", http.StatusInternalServerError)
		return
	}

	if err = tpl.Execute(w, nil); err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d....", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
