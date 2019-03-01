package main

import "os"
import "fmt"

func main() {
        defer fmt.Println("!")

        os.Exit(3)
}
