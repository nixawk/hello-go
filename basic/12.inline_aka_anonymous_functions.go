package main

import "fmt"

// It's good when you use something only one time and so it doesn't need to be
// extracted into an own external function.

func main() {

	// inline aka anonymous function
	add := func(x int, y int) int {
		return x + y
	}

	fmt.Println(add(1, 2)) /* output: 3 */
	fmt.Println(add(2, 3)) /* output: 5 */

	go func() {
		fmt.Println("goroutine") /* never printed */
	}()
}
