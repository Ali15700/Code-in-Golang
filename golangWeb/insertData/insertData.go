package main

import (
	"html/template"
	"net/http"
)

type prodSpec struct {
	Size   string
	Weight float32
	Desc   string
}

type product struct {
	ProdID int
	Name   string
	Cost   string
	Specs  prodSpec
}

var tpl *template.Template
var iphone product

func main() {
	iphone = product{
		ProdID: 01,
		Name:   "IPHONE 14",
		Cost:   "$1099",
		Specs: prodSpec{
			Size:   "100 x 70 x 7 mm",
			Weight: 65,
			Desc:   "High Price Hight Quality",
		},
	}

	tpl, _ = tpl.ParseGlob("*.html")
	http.HandleFunc("/productdetails", productDetailsHandler)
	http.ListenAndServe(":8080", nil)
}

func productDetailsHandler(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "product.html", iphone)
}
