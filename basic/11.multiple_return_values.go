package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	_, c := vals()

	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}

/*
 * Go has built-in support multiple return values.
 * If you only want a subset of the returned values, use the blank identifier _
 */
