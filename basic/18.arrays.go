package main

import "fmt"

func main() {
	var a = []int{0, 1, 2, 3, 4}
	for i := 0; i < len(a); i++ {
		fmt.Printf("i = %d\n", a[i])
	}
}

/*
 * var variable_name [SIZE] variable_type
 *
 * Go programming language provides a data structure called the array,
 * which can store a fixed-size sequential collection of elements of the
 * same type.
 */
