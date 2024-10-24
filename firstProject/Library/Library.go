package Library

import (
	Book "firstProject/Book"
	"firstProject/Storage"
)

type Library struct {
	Storage     Storage.Storage
	GeneratorId func() int
}

func (library *Library) Add(title string, author string, countPages int, year int) {
	id := library.GeneratorId()
	book := Book.Book{ID: id, Title: title, Author: author, CountPages: countPages, Year: year}
	library.Storage.Add(book)
}

func (library *Library) GetBookByTitle(title string) (Book.Book, bool) {

	switch s := library.Storage.(type) {
	case *Storage.StorageSlice:
		for _, book := range s.AllBooks {
			if book.Title == title {
				return book, true
			}
		}
	case *Storage.StorageMap:
		for _, book := range s.AllBooks {
			if book.Title == title {
				return book, true
			}
		}
	}
	return Book.Book{}, false
}

func (library *Library) ReplaceGeneratorId() Library {
	newLibrary := Library{
		Storage:     &Storage.StorageSlice{},
		GeneratorId: library.GeneratorId,
	}
	switch s := library.Storage.(type) {
	case *Storage.StorageSlice:

		for _, book := range s.AllBooks {
			newLibrary.Add(book.Title, book.Author, book.CountPages, book.Year)
		}
	case *Storage.StorageMap:

		for id, book := range s.AllBooks {
			delete(s.AllBooks, id)
			book.ID = library.GeneratorId()
			newLibrary.Add(book.Title, book.Author, book.CountPages, book.Year)
		}
	}
	return newLibrary
}
