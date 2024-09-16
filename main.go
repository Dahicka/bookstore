package main

import (
	"log"
	"net/http"

	"github.com/Dahicka/bookstore/controller"
	db "github.com/Dahicka/bookstore/database"
	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/book", controller.UpdateBook).Methods("PUT")
	myRouter.HandleFunc("/book", controller.AddNewBook).Methods("POST")
	myRouter.HandleFunc("/books", controller.GetBooks)
	myRouter.HandleFunc("/book/{id}", controller.DeleteBook).Methods("DELETE")
	myRouter.HandleFunc("/book/{id}", controller.GetBookById).Methods("GET")

	log.Fatal(http.ListenAndServe("", myRouter))
}

func main() {
	handleRequests()
	db.CloseDB()
}
