package main

import "fmt"

type point struct {
        x, y int
}

func main() {
        
        p := point{1, 2}

        fmt.Printf("%v\n", p)
        fmt.Printf("%+v\n", p)
        fmt.Printf("%#v\n", p)
        fmt.Printf("%T\n", p)
        fmt.Printf("%t\n", true)
        fmt.Printf("%d\n", 123)
        fmt.Printf("%b\n", 14)
        fmt.Printf("%c\n", 44)
        fmt.Printf("%x\n", 44)
        fmt.Printf("%f\n", 78.9)
        fmt.Printf("%e\n", 123400000.0)
        fmt.Printf("%E\n", 123400000.0)
        fmt.Printf("%s\n", "\"string\"")
        fmt.Printf("%q\n", "\"string\"")
        fmt.Printf("%x\n", "hex this")
        fmt.Printf("%p\n", &p)
        fmt.Printf("%6d\n", 12)
        fmt.Printf("%6.2f\n", 1.2)
        fmt.Printf("%6s\n", "s")
}