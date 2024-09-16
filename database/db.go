package db

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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

func SelectAllWithPagination(pageNum int, limitNum int, w http.ResponseWriter) *sql.Rows {
	offset := (pageNum - 1) * limitNum
	query := "SELECT * FROM books LIMIT ? OFFSET ?"
	rows, err := db.Query(query, limitNum, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return rows
}

func SelectById(id string) *sql.Row {
	query := "SELECT * FROM books WHERE id = ?"
	row := db.QueryRow(query, id)
	return row
}

func InsertNewBook(name string, author string, punlished int, w http.ResponseWriter) sql.Result {
	query := "INSERT INTO books (name, author, published) VALUES (?, ?, ?)"
	result, err := db.Exec(query, name, author, punlished)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return result
}

func DeleteById(id string, w http.ResponseWriter) sql.Result {
	query := "DELETE FROM books WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return result
}

func UpdateById(name string, author string, published int, id int, w http.ResponseWriter) {
	query := `UPDATE books SET name = ?, author = ?, published = ? WHERE id = ?`
	_, err := db.Exec(query, name, author, published, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
