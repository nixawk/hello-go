package main

import "fmt"
import "sort"

func main() {

        ints := []int{3, 1, 2}
        sort.Ints(ints)

        fmt.Println("Ints  : ", ints)
        fmt.Println("Sorted: ", sort.IntsAreSorted(ints))
}

/* Go's sort package implements sorting for builtins and user-defined types. */
/* https://gobyexample.com/sorting */