package main

import "fmt"
import "sort"

func main() {

        strs := []string{"c", "a", "b"}
        sort.Strings(strs)

        fmt.Println("Strings: ", strs)
        fmt.Println("Sorted : ", sort.StringsAreSorted(strs))
}

/* Go's sort package implements sorting for builtins and user-defined types. */
/* https://gobyexample.com/sorting */