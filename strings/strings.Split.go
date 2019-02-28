package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
	for i, v := range s.Split("A-B-C-D-E", "-") {
		p(i, v)
	}
}
