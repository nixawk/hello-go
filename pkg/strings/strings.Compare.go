package main

import "fmt"
import "strings"

func main() {
        fmt.Println("Compare: ", strings.Compare("hello", "hello"))  /* Compare:  0 */   
        fmt.Println("Compare: ", strings.Compare("hello", "Hello"))  /* Compare:  1 */    
}

/* func Compare(a, b string) int */
/* https://golang.org/pkg/strings/ */
