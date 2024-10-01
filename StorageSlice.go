package main

type StorageSlice struct {
	allBooks []Book
}

func (books *StorageSlice) Add(book Book) bool {
	books.allBooks = append(books.allBooks, book)
	return true
}

func (books *StorageSlice) GetBookById(id int) (Book, bool) {
	for _, book := range books.allBooks {
		if book.ID == id {
			return book, true
		}
	}
	return Book{}, false
}
