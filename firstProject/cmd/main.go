package main

import (
	Book "firstProject/Book"
	"firstProject/Generator"
	"firstProject/Library"
	"firstProject/Storage"
	"fmt"
)

func printBook(book Book.Book) {
	fmt.Println("Найденная книга: ")
	fmt.Println("ID - ", book.ID)
	fmt.Println("Название - ", book.Title)
	fmt.Println("Автор - ", book.Author)
	fmt.Println("Год выпуска - ", book.Year)
	fmt.Println("Количество страниц - ", book.CountPages)
	fmt.Println()

}

func main() {
	books := []Book.Book{
		{Title: "Дело, которому ты служишь", Author: "Юрий Герман", CountPages: 420, Year: 1958},
		{Title: "Повелитель мух", Author: "Уильям Голдинг", CountPages: 190, Year: 1954},
		{Title: "451 градус по Фаренгейту", Author: "Рэй Бредбери", CountPages: 200, Year: 1953},
	}

	library := Library.Library{
		Storage:     &Storage.StorageSlice{},
		GeneratorId: Generator.FirstGeneratorId(),
	}

	for _, book := range books {
		library.Add(book.Title, book.Author, book.CountPages, book.Year)
	}

	bookFirst, flag := library.GetBookByTitle("Дело, которому ты служишь")
	if flag {
		printBook(bookFirst)
	} else {
		fmt.Println("Книга не найдена")
	}

	bookSecond, flag := library.GetBookByTitle("451 градус по Фаренгейту")
	if flag {
		printBook(bookSecond)
	} else {
		fmt.Println("Книга не найдена")
	}

	library.GeneratorId = Generator.SecondGeneratorId()
	library = library.ReplaceGeneratorId()

	bookThird, flag := library.GetBookByTitle("Повелитель мух")
	if flag {
		printBook(bookThird)
	} else {
		fmt.Println("Книга не найдена")
	}

	library.Storage = &Storage.StorageMap{AllBooks: make(map[int]Book.Book)}

	for _, book := range books {
		library.Add(book.Title, book.Author, book.CountPages, book.Year)
	}

	bookFourth, flag := library.GetBookByTitle("Повелитель мух")
	if flag {
		printBook(bookFourth)
	} else {
		fmt.Println("Книга не найдена")
	}

	bookFifth, flag := library.GetBookByTitle("451 градус по Фаренгейту")
	if flag {
		printBook(bookFifth)
	} else {
		fmt.Println("Книга не найдена")
	}
}
