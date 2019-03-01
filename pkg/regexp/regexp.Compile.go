package main

import "regexp"
import "fmt"

func main() {
        r, _ := regexp.Compile("p[a-z]+ch")
        fmt.Println(r.MatchString("peach"))
}