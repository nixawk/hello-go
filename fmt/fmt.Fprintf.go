package main

import "fmt"
import "os"

func main() {
        fmt.Fprintf(os.Stderr, "this is an %s error\n", "io")
}