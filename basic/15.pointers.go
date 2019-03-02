package main

import "fmt"

func pint(pi *int) {
	fmt.Printf("ptr: %p, val: %d\n", pi, *pi)
}

func main() {
	var i int = 1234

	pint(&i)

	/* Another way to get a pointer is to use the built-in new function */
	p := new(int) /* nil pointer */
	*p = 1111
	pint(p)

	/* In some programming languages there is a significant different
	   between using new and &, with great care being needed to eventally
	   delete anything created with new.
	   Go is not like this, it's garbage collected programming language
	   which means memory is cleaned up automatically when nothing refers
	   to it anymore. */
}

/*
 * Go compiler assign a Nil value to a pointer variable in case you do not
 * have exact address to be assigned. This is done at the time of variable
 * declaration. A pointer that is assigned nil is called a nil pointer.
 */
