package model

type Book struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Published int    `json:"published"`
}

func NewBook(id int, name string, author string, published int) Book {
	return Book{Id: id, Name: name, Author: author, Published: published}
}
