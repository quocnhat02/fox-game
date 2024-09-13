package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"senior/internal/model"

	"github.com/gorilla/mux"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to senior")
}

func HandleBooks(w http.ResponseWriter, r *http.Request) {
	books := model.FetchBooks()
	json.NewEncoder(w).Encode(books)
}

func HandleBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	book := model.FetchBook(id)
	json.NewEncoder(w).Encode(book)
}

func HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	json.NewDecoder(r.Body).Decode(&book)
	books := model.CreateBook(book)
	json.NewEncoder(w).Encode(books)
}

func HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	json.NewDecoder(r.Body).Decode(&book)
	books := model.UpdateBook(book)
	json.NewEncoder(w).Encode(books)
}
