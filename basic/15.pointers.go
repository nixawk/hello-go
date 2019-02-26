package main

import "fmt"

func main() {
        var i int = 1234;
        var p *int

        fmt.Printf("ptr: %p\n", p)  /* nil pointer */
        p = &i
        fmt.Printf("ptr: %p, val: %d\n", p, *p)
}

/*
 * Go compiler assign a Nil value to a pointer variable in case you do not
 * have exact address to be assigned. This is done at the time of variable
 * declaration. A pointer that is assigned nil is called a nil pointer.
 */