package main

import "fmt"
import "strings"

func main() {
        var a = []string{"c", "cpp", "go", "python"}
        fmt.Println(strings.Join(a, ","))
}

/* https://golang.org/pkg/strings/#Join */
