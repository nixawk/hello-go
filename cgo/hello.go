package main

// #include <hello.c>
import "C"
import (
	"fmt"
)

func main() {
	C.hello()
	fmt.Printf("1 + 2 = %d\n", C.add(1, 2))
}

/* https://golang.org/cmd/go/ */
/* https://www.programming-books.io/essential/go/d3b0419346904e0bbf6dc28b69fc93c6-calling-c-from-go-with-cgo */