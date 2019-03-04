package main

import "fmt"

/* Each package has the ability to provide as many init functions as necessary
to be invoked at the beginning of execution time. All the init functions that
are discovered by the compiler are scheduled to be executed prior to the main
function being executed.

The init functions are great for setting up packages, initializing variables, or performing any other bootstrapping you may need prior to the program running. */

func init() {
	fmt.Println("init function")
}

func main() {
	fmt.Println("main function")
}

/*

$ go run 47.init.go
init function
main function

*/
