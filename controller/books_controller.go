package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	db "github.com/Dahicka/bookstore/database"
	"github.com/Dahicka/bookstore/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "5"
	}
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}
	limitNum, err := strconv.Atoi(limit)
	if err != nil || limitNum < 1 {
		http.Error(w, "Invalid limit number", http.StatusBadRequest)
		return
	}
	rows := db.SelectAllWithPagination(pageNum, limitNum, w)
	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.Id, &book.Name, &book.Author, &book.Published); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	row := db.SelectById(id)
	var book model.Book
	if err := row.Scan(&book.Id, &book.Name, &book.Author, &book.Published); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Id does not exist", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	var newBook model.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := db.InsertNewBook(newBook.Name, newBook.Author, newBook.Published, w)

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "New book added:%d, %s, %s, %d", lastInsertID, newBook.Name, newBook.Author, newBook.Published)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result := db.DeleteById(id, w)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Id does not exist", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var newBook model.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.UpdateById(newBook.Name, newBook.Author, newBook.Published, newBook.Id, w)
	w.WriteHeader(http.StatusNoContent)
}
