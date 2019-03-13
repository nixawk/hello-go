package main

import "fmt"

// Creating a new array looks like var array = [size]Type
func create_array() {
	a := [3]int{0, 1, 2}      // classical way
	b := [...]int{0, 1, 2, 3} // a less verbose way
	c := []int{0, 1, 2, 3, 4}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// Array values should be accessed using a number specifying the location of
// the desired value in the array.
// Indexes start at 0 and finish at array length -1.
func index_array() {
	type Person struct {
		name string
		age  int
	}

	ps := [2]Person{{name: "k", age: 20}, {name: "v", age: 25}}
	for _, p := range ps {
		fmt.Printf("name: %s, age: %d\n", p.name, p.age)
	}
}

func main() {
	create_array()
	index_array()
}

/* https://www.programming-books.io/essential/go/f9ebde1b6f7b4ac49a7a4693f0f19a6a-array-indexes */
