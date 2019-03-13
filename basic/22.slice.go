package main

import "fmt"

func create_slice() {
	/* A slice is a segment of an array.
	   Like arrays slices are indexable and have a length.
	   Unlike arrays this length is allowed to change. */

	s1 := make([]int, 5)    /* 5 is the length of slice */
	s2 := make([]int, 5, 8) /* 8 is the capacity of the underlying array */

	fmt.Printf("[create_slice] %v\n", s1)
	fmt.Printf("[create_slice] %v\n", s2)
}

func index_slice() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("[index_slice]")
	for i := range s {
		fmt.Printf(" %d ", i)
	}
	fmt.Printf("\n")
}

func append_slice() {
	s := make([]int, 1)
	fmt.Printf("[append_slice] slice: %v, len: %d, cap: %d\n", s, len(s), cap(s))

	s = append(s, 1)
	fmt.Printf("[append_slice] slice: %v, len: %d, cap: %d\n", s, len(s), cap(s))
}

func copy_slice() {
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)

	fmt.Printf("[copy_slice] src: %#v\n", src)
	fmt.Printf("[copy_slice] dst: %#v\n", dst)
}

func zero_slice() {
	s := make([]int, 5)
	fmt.Printf("[zero_slice] %#v\n", s)

	s = []int(nil)
	fmt.Printf("[zero_slice] %#v\n", s)
}

func slice_length() {
	// The [len()] function returns the elements presents in the slice.
	s := make([]int, 1)
	fmt.Printf("[slice_length] %d\n", len(s)) /* output: 1 */
}

func slice_capacity() {
	// cap() function returns the capacity of the slice.
	s := make([]int, 1, 5)
	fmt.Printf("[slice_capacity] %d\n", cap(s)) /* output: 5 */
}

func main() {
	create_slice()
	index_slice()
	append_slice()
	copy_slice()
	zero_slice()
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
/* https://www.programming-books.io/essential/go/f400553890d34185ba795870807c2615-slices */
