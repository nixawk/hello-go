package main

import "fmt"

type Request struct {
	Resource string
}

type AuthenticatedRequest struct {
	Request            /* Request is an embedded field. */
	Username, Password string
}

func main() {
	ar := new(AuthenticatedRequest)
	ar.Resource = "example.com/request"
	ar.Username = "Sam"
	ar.Password = "P@ssword"
	fmt.Printf("%#v\n", ar)
}

// Composition provides an alternative to inheritance. A struct may include
// another type by name in its declaration;
/* https://www.programming-books.io/essential/go/5c99711b5d2b467d85fe082b1bef3268-composition-and-embedding */
