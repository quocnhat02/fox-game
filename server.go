package main

import (
	"fmt"
	"net/http"
	"os"
)

func bookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		genre := r.URL.Query().Get("genre")
		title := r.URL.Query().Get("title")
		w.WriteHeader(http.StatusOK) 
		w.Write([]byte(fmt.Sprintf("Book retrieved. Genre: %s, Title: %s", genre, title)))
	case "POST":
		genre := r.Body.Genre
		title := r.Body.Title
		if genre == "" || title == "" {
			http.Error(w, "Genre and title are required", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf("Book created. Genre: %s, Title: %s", genre, title)))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! I am a server.")
	})

	http.HandleFunc("/books", bookHandler)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
