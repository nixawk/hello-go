package main

import "fmt"

func main() {
        var numbers []int            /* a slice of unspecified size */

        fmt.Printf("cap=%d, slice=%v\n", cap(numbers), numbers) /* slice default, nil */

        numbers = make([]int, 5, 5)  /* a slice of length 5 and capacity 5 */

        for i := 0; i < len(numbers); i++ {
                fmt.Printf("numbers[%d] = %d\n", i, numbers[i])
        }

        numbers = append(numbers, 1)

        fmt.Printf("cap=%d, slice=%v\n", cap(numbers), numbers)  /* the capacity of the slice */
}

/*
 * To define a slice, you can delare it as an array without specifying its
 * size. Alternatively, you can use [make] function to create a slice.
 */


/*
 * The [len()] function returns the elements presents in the slice where
 * cap() function returns the capacity of the slice.
 */

/*
 * If a slice is declared with no inputs, then by default, it is initialized
 * as nil. Its length and capacity are zero.
 */

/* https://www.tutorialspoint.com/go/go_slice.htm */
