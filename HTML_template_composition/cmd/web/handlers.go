package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	//Initialize a slice containing the paths to the two files. It's important
	//to note that file containing our base template must be the  *first*

	//file in slice.

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	//Use the template.ParseFiles() function to read the files and store the
	//templates in a template set. Notice that we can pass the slice of file
	//path as a variadic parameters?

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	//Use the ExecuteTemplate() method to write the content of the "base"
	//template as teh response body.

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
