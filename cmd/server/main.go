package main

import (
	"fmt"
	"net/http"
	"os"
	"senior/internal/api"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", api.HandleIndex).Methods("GET")
	r.HandleFunc("/books", api.HandleBooks).Methods("GET")
	r.HandleFunc("/books/{id}", api.HandleBook).Methods("GET")
	r.HandleFunc("/books", api.HandleCreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", api.HandleUpdateBook).Methods("PUT")
 
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
