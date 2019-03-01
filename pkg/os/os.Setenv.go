package main

import "os"
import "fmt"

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func main() {
        fmt.Printf("SHELL=%s\n", os.Getenv("SHELL"))

        err := os.Setenv("SHELL", "/bin/sh")
        check(err)

        fmt.Printf("SHELL=%s\n", os.Getenv("SHELL"))
}
