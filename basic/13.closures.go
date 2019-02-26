package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()

	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", nextInt())
	}
}

/*
 * Go supports anonymous functions, which can form closures. Anonymous
 * functions are useful when you want to define a function inline without
 * having to name it.
 */

/* https://gobyexample.com/closures */
