package main

import "fmt"

/* Like a struct an interface is created using the type keyword, followed by
   a name and the keyword interface. But instead of defining fields, we
   defines a "method set" */

type Inter interface {
	get_name() string
}

type Person struct {
	name string
	age  int
}

type Fruit struct {
	name string
	loc  string
}

func (p Person) get_name() string {
	return p.name
}

func (f Fruit) get_name() string {
	return f.name
}

func print_name(i Inter) {
	fmt.Printf("my name is %s\n", i.get_name())
}

func main() {
	p := Person{name: "mutai", age: 25}
	f := Fruit{name: "peach", loc: "jp"}

	print_name(p)
	print_name(f)
}

/*
 * Go programming provides another data type called interfaces which represents
 * a set of method signatures.
 */
