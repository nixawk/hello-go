package main

import "fmt"

func main() {
	m := map[int]string{
		0: "c",
		1: "cpp",
		2: "golang",
	}

	for key, value := range m {
		fmt.Printf("key: %d, value: %s\n", key, value)
	}
}

/* https://www.programming-books.io/essential/go/c522a62872884110bcc300db67e0e9ad-range-statement#1 */
