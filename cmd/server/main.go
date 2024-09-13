package main

import (
	"fmt"
	"net/http"
	"os"
	"senior/internal/api"
)

func main() {
	mux := http.NewServeMux() 

	mux.HandleFunc("/", api.HandleIndex)
	mux.HandleFunc("/books", api.HandleBooks)
 
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
