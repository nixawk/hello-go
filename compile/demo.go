package main

import "fmt"

func main() {
	fmt.Println("This is a demo for cross compiling")
}

/*

Cross-compiling works with environment variables. If you do `go build` go
automatically uses your operating system and architecture to complete the
binary.

If you now want to compile it to run on a mac 64bit system you could do
`GOOS=darwin GOARCH=amd64 go build` and go will produce a binary that
can be run on a Mac but not on my machine.

If you want to name the binary accordingly you could do something like:
`GOOS=darwin GOARCH=amd64 go build -o darwin-binary`

*/
