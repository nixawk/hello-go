package main

import "fmt"

func main() {

	/* A slice is a segment of an array.
	   Like arrays slices are indexable and have a length.
	   Unlike arrays this length is allowed to change. */

	/* make */
	s1 := make([]int, 5)    /* 5 is the length of slice */
	s2 := make([]int, 5, 8) /* 8 is the capacity of the underlying array */

	fmt.Println(s1, len(s1), cap(s1)) /* [0 0 0 0 0] 5 5 */
	fmt.Println(s2, len(s2), cap(s2)) /* [0 0 0 0 0] 5 8 */

	/* append */
	sa1 := append(s1, 1, 2, 3, 4)
	sa2 := append(s2, 1, 2, 3, 4)

	fmt.Println(sa1, len(sa1), cap(sa1)) /* [0 0 0 0 0 1 2 3 4] 9 10 */
	fmt.Println(sa2, len(sa2), cap(sa2)) /* [0 0 0 0 0 1 2 3 4] 9 16 */

	/* copy */
	sc1 := []int{1, 2, 3}
	sc2 := make([]int, 2)

	copy(sc2, sc1) /* copy(dst, src) */

	fmt.Println(sc1, sc2) /* [1 2 3] [1 2] */
}

/*
 * To define a slice, you can delare it as an array without specifying its
 * size. Alternatively, you can use [make] function to create a slice.
 */

/*
 * The [len()] function returns the elements presents in the slice where
 * cap() function returns the capacity of the slice.
 */

/*
 * If a slice is declared with no inputs, then by default, it is initialized
 * as nil. Its length and capacity are zero.
 */

/* https://www.tutorialspoint.com/go/go_slice.htm */
