package main

type Library struct {
	storage     Storage
	generatorId func() int
}

func (library *Library) Add(title string, author string, countPages int, year int) {
	id := library.generatorId()
	book := Book{ID: id, Title: title, Author: author, CountPages: countPages, Year: year}
	library.storage.Add(book)
}

func (library *Library) GetBookByTitle(title string) (Book, bool) {

	switch s := library.storage.(type) {
	case *StorageSlice:
		for _, book := range s.allBooks {
			if book.Title == title {
				return book, true
			}
		}
	case *StorageMap:
		for _, book := range s.allBooks {
			if book.Title == title {
				return book, true
			}
		}
	}
	return Book{}, false
}

func (library *Library) ReplaceGeneratorId() Library {
	newLibrary := Library{
		storage:     &StorageSlice{},
		generatorId: library.generatorId,
	}
	switch s := library.storage.(type) {
	case *StorageSlice:

		for _, book := range s.allBooks {
			newLibrary.Add(book.Title, book.Author, book.CountPages, book.Year)
		}
	case *StorageMap:

		for id, book := range s.allBooks {
			delete(s.allBooks, id)
			book.ID = library.generatorId()
			newLibrary.Add(book.Title, book.Author, book.CountPages, book.Year)
		}
	}
	return newLibrary
}
