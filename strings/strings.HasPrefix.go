package main

import "strings"
import "fmt"

func main() {
	fmt.Println(strings.HasPrefix("hello world", "he"))
	fmt.Println(strings.HasPrefix("hello world", "go"))
}
