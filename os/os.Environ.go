package main

import "os"
import "fmt"
import "strings"

func main() {
        for _, e := range os.Environ() {
                p := strings.Split(e, "=")
                fmt.Printf("{'key': '%s', 'val': '%s'}\n", p[0], p[1])
        }
}

/* https://gobyexample.com/environment-variables */
