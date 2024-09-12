package main

import "fmt"

func main() {
	// Declaring and initializing variables
	var name string = "World"
	age := 42

	// Using variables in output
	fmt.Println("Hello,", name)
	fmt.Printf("You are %d years old\n", age)

	// Modifying variables
	name = "Gopher"
	age++

	// Using modified variables
	fmt.Printf("Hello, %s! Next year, you'll be %d\n", name, age)
}