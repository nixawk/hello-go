package main

import "regexp"
import "fmt"

func main() {
        match, _ := regexp.MatchString("p[a-z]+ch", "peach")
        fmt.Println(match)
}