package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
	p(s.Replace("hello world ", " ", ",", 1))  /* output: hello,world */
	p(s.Replace("hello world ", " ", "!", -1)) /* output: hello!world! */
}
