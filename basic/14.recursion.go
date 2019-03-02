package main

import "fmt"

func algorithm(x int) int {
	var y int

	y = x * 2
	if y > 100 || y < -100 {
		return y
	}

	return algorithm(y)
}

func main() {
	var n int

	fmt.Printf("[*] enter a number: ")
	fmt.Scanf("%d", &n)
	fmt.Printf("%d\n", algorithm(n))
}

/*
 * Recursion is the process of repeating items in a self-similar way.
 */
