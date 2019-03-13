package main

import "fmt"

func create_slice() {
	/* A slice is a segment of an array.
	   Like arrays slices are indexable and have a length.
	   Unlike arrays this length is allowed to change. */

	s1 := make([]int, 5)    /* 5 is the length of slice */
	s2 := make([]int, 5, 8) /* 8 is the capacity of the underlying array */

	fmt.Println(s1)
	fmt.Println(s2)
}

func index_slice() {
	s := []int{0, 1, 2, 3, 4, 5}
	for i := range s {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")
}

func append_slice() {
	s := make([]int, 1)
	fmt.Printf("[B]slice: %v, len: %d, cap: %d\n", s, len(s), cap(s))

	s = append(s, 1)
	fmt.Printf("[A]slice: %v, len: %d, cap: %d\n", s, len(s), cap(s))
}

func slice_length() {
	// The [len()] function returns the elements presents in the slice.
	s := make([]int, 1)
	fmt.Println(len(s)) /* output: 1 */
}

func slice_capacity() {
	// cap() function returns the capacity of the slice.
	s := make([]int, 1, 5)
	fmt.Println(cap(s)) /* output: 5 */
}

func main() {
	create_slice()
	index_slice()
	append_slice()
	slice_length()
	slice_capacity()
}

/*
 * Slice is a growable sequence of values of the same type.
 * Slice has length and capacity.
 * Capacity represents how many total elements a slice can have. That's the size of underlying array.
 * Length is the current number of elements in the slice.
 *
 * If a slice is declared with no inputs, then by default, it is initialized
 * as nil. Its length and capacity are zero.
 */

/* https://www.tutorialspoint.com/go/go_slice.htm */
