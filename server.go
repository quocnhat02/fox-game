package main

import (
	"fmt"
	"net/http"
)

type MessageHandler struct {
	Message string
}

func (m *MessageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.Message)
}

func main() {
	messageHandler := &MessageHandler{Message: "Hello, World! I am a server."}

	http.Handle("/", messageHandler)
	http.ListenAndServe(":8080", nil)
}
