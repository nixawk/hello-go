package main

import "fmt"

type Account struct {
	Name string
}

// Since SetName() mutates the instance, the receiver must be a pointer in
// order to effect a permanent change in the instance.
func (a *Account) SetName(n string) {
	a.Name = n
}

func (a *Account) GetName() string {
	return a.Name
}

func main() {
	a := Account{Name: "King"}
	fmt.Printf("Name: %s\n", a.Name)
	a.SetName("Queen")
	fmt.Printf("Name: %s\n", a.GetName())
}

/*

$ go run 16.structures_methods.go
Name: King
Name: Queen

*/
