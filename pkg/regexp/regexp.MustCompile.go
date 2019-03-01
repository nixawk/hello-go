package main

import "regexp"
import "fmt"

func main() {
        r := regexp.MustCompile("p[a-z]+ch")
        fmt.Println(r)
}
