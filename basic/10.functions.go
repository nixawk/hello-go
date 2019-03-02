package main

import "fmt"

/*
func function_name( [parameter list] ) [return_types]
{
   body of the function
}
*/

/* Functions are also known as method, sub-routine, or procedure */

func add(x, y int) int {
	var z int

	z = x + y
	fmt.Printf("x + y = %d\n", z)
	return z
}

func foo() (string, string) {
	return "Go", "Python"
}

func main() {
	add(1, 2)

	s1, s2 := foo()
	fmt.Printf("%s VS %s\n", s1, s2)
}
