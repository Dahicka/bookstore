package main

import (
	"log"
	"net/http"

	"github.com/Dahicka/bookstore/controller"
	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/book", controller.AddNewBook).Methods("POST")
	myRouter.HandleFunc("/book", controller.UpdateBook).Methods("PUT")
	myRouter.HandleFunc("/books", controller.GetBooks)
	myRouter.HandleFunc("/book/{id}", controller.DeleteBook).Methods("DELETE")
	myRouter.HandleFunc("/book/{id}", controller.ReturnBook)

	log.Fatal(http.ListenAndServe("", myRouter))
}

func main() {
	handleRequests()
	controller.CloseDB()
}
