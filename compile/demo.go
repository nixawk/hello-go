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

/*

$ go tool dist list
android/386
android/amd64
android/arm
android/arm64
darwin/386
darwin/amd64
darwin/arm
darwin/arm64
dragonfly/amd64
freebsd/386
freebsd/amd64
freebsd/arm
js/wasm
linux/386
linux/amd64
linux/arm
linux/arm64
linux/mips
linux/mips64
linux/mips64le
linux/mipsle
linux/ppc64
linux/ppc64le
linux/riscv64
linux/s390x
nacl/386
nacl/amd64p32
nacl/arm
netbsd/386
netbsd/amd64
netbsd/arm
openbsd/386
openbsd/amd64
openbsd/arm
plan9/386
plan9/amd64
plan9/arm
solaris/amd64
windows/386
windows/amd64

*/
