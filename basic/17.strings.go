package main

import "fmt"
import "strings"

func foo_strlen() {
        var s string = "Hello, Go"
        fmt.Printf("The length of [%s] is %d\n", s, len(s))     
}

func foo_concat() {
        s := []string{"Hello", "world!"}
        fmt.Println(strings.Join(s, " "))
}

func main() {

        fmt.Println("This is a demo string.")
        fmt.Println(`This is a demo string.`)

        foo_strlen()
        foo_concat()
}

/*
 * Strings, which are widely used in Go programming, are a readonly slice of
 * bytes. In the Go programming language, strings are slices. The Go platform
 * provides various libraries to manipulate strings.
 *
 * - unicode
 * - regexp
 * - strings
 *
 */

/*
 * len(str) method returns the number of bytes contained in the string literal.
 */
