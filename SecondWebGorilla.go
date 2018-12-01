package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

  bookrouter := router.PathPrefix("/bks").Subrouter()
  bookrouter.HandleFunc("/", AllBooks)
  bookrouter.HandleFunc("/{title}", GetBook)

  router.HandleFunc("/books/{title}", CreateBook).Methods("POST")
  router.HandleFunc("/books/{title}", ReadBook).Methods("GET")
  router.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
  router.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")
  router.HandleFunc("/booksx/{title}", BookHandler).Host("localhost")
  router.HandleFunc("/secure", SecureHandler).Schemes("https")
  router.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	http.ListenAndServe(":80", router )
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  title := vars["title"]
  fmt.Fprintf(w, "CreateBook You've requested the book: %s \n", title)
}
func ReadBook(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  title := vars["title"]
  fmt.Fprintf(w, "ReadBook You've requested the book: %s \n", title)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  title := vars["title"]
  fmt.Fprintf(w, "UpdateBook You've requested the book: %s \n", title)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  title := vars["title"]
  fmt.Fprintf(w, "DeleteBook You've requested the book: %s \n", title)
}
func BookHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  title := vars["title"]
  fmt.Fprintf(w, "BookHandler You've requested the book: %s \n", title)
}
func SecureHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "SecureHandler")
}
func InsecureHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "InsecureHandler ")
}
func AllBooks(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "AllBooks ")
}
func GetBook(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "GetBook ")
}
