package main

import "fmt"

func panic_recover() {
        str := recover()
        fmt.Println(str)
}

func main() {
        defer panic_recover()
        panic("PANIC")
        fmt.Println("after panic recover") /* never printed */
}

/* Earlier we created a function that called the panic function to cause a
   run time error. We can handle a run-time panic with the built-in recover
   function. recover stops the panic and returns the value that was passed
   to the call to panic. */

/*

$ go run recover.go 
PANIC

*/