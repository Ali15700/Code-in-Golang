package main

import (
	"html/template"
	"net/http"
)

type GroceryList []string

var tpl *template.Template
var g GroceryList

func main() {
	g = GroceryList{"Oil", "butter", "chica beans", "cheese"}
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/list", listHandler)

	http.ListenAndServe(":8080", nil)
}

func listHandler(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "grocery.html", g)
}
