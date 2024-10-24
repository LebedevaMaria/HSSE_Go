package Storage

import Book "firstProject/Book"

type Storage interface {
	GetBookById(id int) (Book.Book, bool)
	Add(book Book.Book) bool
}
