package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Dahicka/bookstore/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:test123@tcp(127.0.0.1:3306)/booksdb")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}
func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Println("Error closing DB:", err)
		}
	}
}

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

	offset := (pageNum - 1) * limitNum
	query := "SELECT * FROM books LIMIT ? OFFSET ?"
	rows, err := db.Query(query, limitNum, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

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
	query := "SELECT * FROM books WHERE id = ?"
	row := db.QueryRow(query, id)
	var book model.Book
	if err := row.Scan(&book.Id, &book.Name, &book.Author, &book.Published); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "ID does not exist", http.StatusNotFound)
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

	query := "INSERT INTO books (name, author, published) VALUES (?, ?, ?)"
	_, err = db.Exec(query, newBook.Name, newBook.Author, newBook.Published)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var lastInsertID int64
	selectQuery := "SELECT LAST_INSERT_ID()"
	err = db.QueryRow(selectQuery).Scan(&lastInsertID)
	if err != nil {
		log.Println("Error retrieving last insert ID:", err)
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

	query := "DELETE FROM books WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Id does not exist", http.StatusNotFound)
		return
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var newBook model.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE books SET name = ?, author = ?, published = ? WHERE id = ?`
	_, err = db.Exec(query, newBook.Name, newBook.Author, newBook.Published, newBook.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
