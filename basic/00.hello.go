package main

import "fmt"

/* All init function in any code that are part of the program will get called
   before the main function */
func init() {
	fmt.Println("I'm the init func.")
}

func main() {
	fmt.Println("Hello, World !")
}

/* go run hello.go */
/* go build hello.go && ./hello */

/* $ go run 00.hello.go
   I'm the init func.
   Hello, World !
*/

/*
 * The first line of the program package main defines the package name in which
 * this program should lie. It is a mandatory statement, as Go programs run in
 * packages. The main package is the starting point to run the program. Each
 * package has a path and name associated with it.

 */
