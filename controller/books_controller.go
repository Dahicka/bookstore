package controller

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, %s id", id)
}