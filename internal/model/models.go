package model

type Book struct {
	Title  string
	Author string
	Pages  int
}

func FetchBooks() []Book {
	return []Book{
		{
			Title:  "The Go Programming Language",
			Author: "Alan A. A. Donovan",
			Pages:  380,
		},
		{
			Title:  "Go in Action",
			Author: "William Kennedy",
			Pages:  300,
		},
	}
}