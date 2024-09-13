package model

type Book struct {
	ID     string
	Title  string
	Author string
}

func FetchBooks() []Book {
	return []Book{
		{
			ID:     "1",
			Title:  "The Go Programming Language",
			Author: "Alan A. A. Donovan",
		},
		{
			ID:     "2",
			Title:  "Go in Action",
			Author: "William Kennedy",
		},
		{
			ID:     "3",
			Title:  "The Go Programming Language",
			Author: "Alan A. A. Donovan",
		},
	}
}

func FetchBook(id string) Book {
	books := FetchBooks()
	for _, book := range books {
		if book.ID == id {
			return book
		}
	}
	return Book{}
}

func CreateBook(book Book) []Book {
	books := FetchBooks()
	books = append(books, book)
	return books
}

func UpdateBook(book Book) []Book {
	books := FetchBooks()
	for i, b := range books {
		if b.ID == book.ID {
			books[i] = book
			return books
		}
	}
	return books
}
