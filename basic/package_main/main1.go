package main

import "fmt"

func main() {
	fmt.Println("main1.go --- main")
}

/*

$ go run *.go
# command-line-arguments
./main2.go:5:6: main redeclared in this block
	previous declaration at ./main1.go:5:6

*/