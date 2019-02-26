package main


import "fmt"


func main() {
        var my_str string  /* Static Type Declaration */
        var i, j int

        my_str = "Hello, Go"
        fmt.Println(my_str)

        i, j = 2, 3        /* Mixed Variable Declaration */
        my_int := i * j    /* Dynamic Type Declaration */
        fmt.Println(my_int)
}


/*
 * Boolean types: true, false
 * Numeric types: integer, floating
 *                uint8 / uint16 / uint32 / uint64
 *                int8 / int16 / int32 / int64
 *                float32 / float64
 *                complex64 / complex128
 *                byte (same as uint8)
 *                rune (same as int32)
 *                uint (32 or 64 bits)
 *                int  (same size as uint)
 *                uintptr (an unsigned integer pointer)
 * String types: 
 * Derived types: Pointer, Array, Structure, Union, Function,
                  Slice, Inteface, Map, Channel
 */

/*
 * Variable Definition in Go
 *
 * A variable definition tells the compiler where and how much storage to
 * create for the variable.
 *
 *      var variable_list optional_data_type
 *
 */

/*
 * Static Type Declaration
 *
 * For definition without an intializer: variables with static storage
 * duration are implcitly initialized with nil (all bytes have the value 0);
 * the initial value of all other variables is zero value of their data type.
 */

/*
 * Dynamic Type Declaration
 *
 * A dynamic type variable declaration requires the compiler to interpret
 * the type of the variable based on the value passed to it. The compiler
 * does not require a variable to have type staticlly as a necessary
 * requirement.
 *
 */

 /* https://gobyexample.com/variables */
 