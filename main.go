package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Dahicka/bookstore/controller"
)

func main() {
	http.HandleFunc("/book", controller.HelloHandler)
	fmt.Println("server running")
	err := http.ListenAndServe("", nil)
    if err != nil {
        log.Fatal(err)
    }
}