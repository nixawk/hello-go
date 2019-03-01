package main

import "regexp"
import "fmt"

func main() {
        r, _ := regexp.Compile("p[a-z]+ch")
        fmt.Println(r.FindAllStringIndex("peach punch", -1)) /* output: [[0 5] [6 11]] */
}
