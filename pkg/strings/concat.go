package main

import (
	"bytes"
	"fmt"
)

func concat(s1 string, s2 string) string {
	buffer := new(bytes.Buffer)

	buffer.WriteString(s1)
	buffer.WriteString(s2)

	return buffer.String()
}

func main() {
	fmt.Println(concat("hello, ", "golang"))
}

/* https://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go */
