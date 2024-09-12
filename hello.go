package main

import "fmt"

func printMessage(done chan bool) {
	fmt.Println("Hello, World!")
	done <- true
}

func main() {
	done := make(chan bool)
	go printMessage(done)
	<-done
	fmt.Println("Done")
}