package api

import (
	"fmt"
	"net/http"
	"senior/internal/model"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to senior")
}

func HandleBooks(w http.ResponseWriter, r *http.Request) {
	books := model.FetchBooks()
	for _, book := range books {
		fmt.Fprintf(w, "Book: %s\n", book.Title)
	}
}