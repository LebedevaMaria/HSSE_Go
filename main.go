package main

import "fmt"

func printBook(book Book) {
	fmt.Println("Найденная книга: ")
	fmt.Println("ID - ", book.ID)
	fmt.Println("Название - ", book.Title)
	fmt.Println("Автор - ", book.Author)
	fmt.Println("Год выпуска - ", book.Year)
	fmt.Println("Количество страниц - ", book.CountPages)
	fmt.Println()

}

func main() {
	books := []Book{
		{Title: "Дело, которому ты служишь", Author: "Юрий Герман", CountPages: 420, Year: 1958},
		{Title: "Повелитель мух", Author: "Уильям Голдинг", CountPages: 190, Year: 1954},
		{Title: "451 градус по Фаренгейту", Author: "Рэй Бредбери", CountPages: 200, Year: 1953},
	}

	library := Library{
		storage:     &StorageSlice{},
		generatorId: FirstGeneratorId(),
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

	library.generatorId = SecondGeneratorId()
	library = library.ReplaceGeneratorId()

	bookThird, flag := library.GetBookByTitle("Повелитель мух")
	if flag {
		printBook(bookThird)
	} else {
		fmt.Println("Книга не найдена")
	}

	library.storage = &StorageMap{allBooks: make(map[int]Book)}

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

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
