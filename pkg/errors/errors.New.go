package main

import "fmt"
import "errors"

func panic_recover() {
	msg := recover()
	fmt.Println(msg)
}

func main() {
	err := errors.New("this is a custom error")
	defer panic_recover()
	panic(err)
}
