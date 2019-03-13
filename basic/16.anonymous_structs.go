package main

import "fmt"

func main() {
	// An anonymous struct doesn't have a name.
	data := struct {
		Number int
		Text   string
	}{
		42,
		"Hello, World",
	}
	fmt.Printf("data: %+v\n", data)
}

/* https://www.programming-books.io/essential/go/4eb5f7b0ca13495997e624a6639d3eea-anonymous-structs */
