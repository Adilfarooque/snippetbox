package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/partials/nav.tmpl",
	}

	ts , err := template.ParseFiles(files...)
	if err != nil{
		log.Print(err.Error())
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}

	if err := ts.ExecuteTemplate(w,"base",nil); err != nil{
		log.Print(err.Error())
		http.Error(w,"Internal Server Error",http.StatusInternalServerError)
	}

}
