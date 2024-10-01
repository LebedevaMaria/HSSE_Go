package main

type StorageMap struct {
	allBooks map[int]Book
}

func (books *StorageMap) Add(book Book) bool {
	books.allBooks[book.ID] = book
	return true
}
func (books *StorageMap) GetBookById(id int) (Book, bool) {
	book, ok := books.allBooks[id]
	return book, ok
}
