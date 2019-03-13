package main

import "fmt"

func main() {
	s := []string{"c", "cpp", "golang"}

	for index, value := range s {
		fmt.Printf("index: %d, value: %s\n", index, value)
	}
}

/* https://www.programming-books.io/essential/go/c522a62872884110bcc300db67e0e9ad-range-statement#1 */
