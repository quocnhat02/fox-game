package main

import "fmt"

func main() {
	fruits := map[string]int{
		"apple":  2,
		"banana": 3,
		"cherry": 1,
	}

	for key, value := range fruits {
		fmt.Printf("%s: %d\n", key, value)
	}
}