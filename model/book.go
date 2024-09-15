package model

import "time"

type Book struct {
	id        int `json:"id"`
	name      string `json:"name"`
	author    string `json:"author"`
	published time.Time `json:"published"`
}

func NewBook(id int, name string, author string, published time.Time) Book {
	return Book{id: id, name: name, author: author, published: published}
}
