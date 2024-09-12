package main

import (
	"fmt"
	"net/http"
	"os"
)

func bookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Book retrieved."))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Book created."))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
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
