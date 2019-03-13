package main

import "fmt"

func create_map() {
	// Map is an unordered collection of key-value pairs.

	// Maps are reference types, and once defined they have a zero value of
	// nil. Writes to nil maps will cause a panic and reads will always
	// return the zero value
	var m1 map[int]string

	// To initialize a map, use the make function.
	m2 := make(map[int]string)

	// Set map initial value with curly brackets.
	m3 := map[int]string{0: "c", 1: "cpp", 2: "golang"}

	fmt.Printf("[create_map] %v, %d\n", m1, len(m1))
	fmt.Printf("[create_map] %v, %d\n", m2, len(m2))
	fmt.Printf("[create_map] %v, %d\n", m3, len(m3))
}

func index_map() {
	m := make(map[int]string)
	m[0] = "c"
	m[1] = "cpp"
	m[2] = "golang"

	for k, v := range m {
		fmt.Printf("[index_map] %d = %s\n", k, v)
	}
}

func len_map() {
	m := map[int]string{0: "c", 1: "cpp", 2: "golang"}
	fmt.Printf("[len_map] %d\n", len(m))
}

func copy_map() {
	src := map[int]string{0: "c", 1: "cpp", 2: "golang"}
	dst := make(map[int]string)

	for k, v := range src {
		dst[k] = v
	}

	fmt.Printf("[copy_map] src: %v\n", src)
	fmt.Printf("[copy_map] dst: %v\n", dst)
}

func del_map() {
	m := map[int]string{0: "c", 1: "cpp", 2: "golang"}
	delete(m, 1)

	fmt.Printf("[del_map] %v\n", m)
}

func main() {
	create_map()
	index_map()
	len_map()
	copy_map()
	del_map()
}

/* https://www.programming-books.io/essential/go/80fb91dd63d445e28010b9f5e261da81-maps */
