package main

import (
	"html/template"
	"net/http"
)

//  {{/* a comment */}}	Defines a comment
/*
{{.}}	Renders the root element
{{.Name}}	Renders the “Name”-field in a nested element
{{if .Done}} {{else}} {{end}}	Defines an if/else-Statement
{{range .List}} {{.}} {{end}}	Loops over all “List” field and renders each using {{.}}
*/

type task struct {
	Name string
	Done bool
}

var tpl *template.Template

var Todo []task

func main() {
	Todo = []task{{"give dog a bath", true}, {"mow the lawn", false}, {"pickup groceries", false}, {"take out trash", false}, {"paint kitchen", false}, {"feed dog", false}, {"water plants", false}}
	tpl, _ = tpl.ParseGlob("*.html")
	http.HandleFunc("/todolist", todoHandler)
	http.ListenAndServe(":8080", nil)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	tpl.ExecuteTemplate(w, "todolist.html", Todo)
}
