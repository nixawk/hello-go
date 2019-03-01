package main

import "fmt"
import "strings"

func main() {
        fmt.Println("ContainsAny: ", strings.ContainsAny("hello", "abc")) /* false */
        fmt.Println("ContainsAny: ", strings.ContainsAny("hello", "olh")) /* true */
}

/* func ContainsAny(s, chars string) bool */
/* https://golang.org/pkg/strings/ */
