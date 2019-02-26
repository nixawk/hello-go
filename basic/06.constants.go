package main

import "fmt"
import "math"

func main() {
        fmt.Println(85)      /* decimal */
        fmt.Println(0125)    /* octal */
        fmt.Println(0x55)    /* hexadecimal */

        var my_str string = "A\tB"
        fmt.Println(my_str)  /* string */

        const n = 500000000  /* const declares a constant value. */
        fmt.Println(math.Sin(n))
}

/*
 * Constants refer to fixed values that the program may not alter during its
 * execution. These fixed values are also called literals.
 *
 *    Integer Literals
 *    Floating Literals
 *    String Literals
 *    Escape Sequence
 */

/* https://gobyexample.com/constants */
