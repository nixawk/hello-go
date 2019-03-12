package main

import "fmt"

// Functions can have named return values, I've never used it but I think it
// comes in handy when you have many variables with slightly similar naming.

func name_return_values() (x int) {
	x = 3
	return
}

func main() {
	fmt.Println("number is :", name_return_values())
}

/*
$ go run named_return_values.go
number is : 3
*/
