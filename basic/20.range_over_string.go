package main

import "fmt"

func main() {
	s := "Hey 世界"
	for index, rune := range s {
		fmt.Printf("index: %d, rune: %d\n", index, rune)
	}
}

/* https://www.programming-books.io/essential/go/c522a62872884110bcc300db67e0e9ad-range-statement#1 */
