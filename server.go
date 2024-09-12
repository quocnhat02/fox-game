package main

import (
	"fmt"
	"net/http"
	"os"
)

func bookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here are the books")
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
