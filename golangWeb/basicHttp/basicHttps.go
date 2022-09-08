package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", homeHandleFunc)
	http.ListenAndServe(":8080", nil)
}

func homeHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Home!!")
}
