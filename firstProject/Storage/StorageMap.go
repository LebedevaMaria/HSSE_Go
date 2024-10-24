package Storage

import Book "firstProject/Book"

type StorageMap struct {
	AllBooks map[int]Book.Book
}

func (books *StorageMap) Add(book Book.Book) bool {
	books.AllBooks[book.ID] = book
	return true
}
func (books *StorageMap) GetBookById(id int) (Book.Book, bool) {
	book, ok := books.AllBooks[id]
	return book, ok
}
