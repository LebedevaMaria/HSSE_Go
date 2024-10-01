package main

type Storage interface {
	GetBookById(id int) (Book, bool)
	Add(book Book) bool
}
