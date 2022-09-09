package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {

	tpl, _ = template.ParseFiles("info.html")

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tpl.Execute(w, nil)
}
