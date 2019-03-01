package main

import "fmt"
import s "strings"

var p = fmt.Println

func main() {
        p("Contains: ", s.Contains("helloworld", "hello"))
        /* Contains:  true */
}

/* func Contains(s, substr string) bool */
/* https://golang.org/pkg/strings/ */
