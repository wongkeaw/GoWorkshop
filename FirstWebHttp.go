package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})
	http.HandleFunc("/first", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to First")
	})
	http.HandleFunc("/second", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Second")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
