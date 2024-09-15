package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Dahicka/bookstore/model"
	"github.com/gorilla/mux"
)

//ReadAllBooks
func GetBooks(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	// w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "id je %s", id)

	book := model.NewBook(1, "abc", "abc", time.Now())
	fmt.Fprint(w, book)
}
//ReadBookById
func ReturnBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    key := vars["id"]

    fmt.Fprintf(w, "Key: " + key)
}

//AddNewBook
func AddNewBook(w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
}

//Delete
func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "deleted with %s id", id)
}

//Update
func UpdateBook(w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
}
