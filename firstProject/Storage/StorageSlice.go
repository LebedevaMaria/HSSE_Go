package Storage

import Book "firstProject/Book"

type StorageSlice struct {
	AllBooks []Book.Book
}

func (books *StorageSlice) Add(book Book.Book) bool {
	books.AllBooks = append(books.AllBooks, book)
	return true
}

func (books *StorageSlice) GetBookById(id int) (Book.Book, bool) {
	for _, book := range books.AllBooks {
		if book.ID == id {
			return book, true
		}
	}
	return Book.Book{}, false
}
