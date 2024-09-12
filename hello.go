package main

import "fmt"

func main() {
	role := "admin"

	switch role {
	case "admin":
		fmt.Println("You are an admin")
	case "user":
		fmt.Println("You are a user")
	default:
		fmt.Println("You are a guest")
	}
}