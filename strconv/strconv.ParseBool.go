package main

import "strconv"
import "fmt"

func main() {
        fmt.Println(strconv.ParseBool("true")) /* true <nil> */
        fmt.Println(strconv.ParseBool("True")) /* true <nil> */
        fmt.Println(strconv.ParseBool("TRUE")) /* true <nil> */

        fmt.Println(strconv.ParseBool("FALSE")) /* false <nil> */

        /* false strconv.ParseBool: parsing "nil": invalid syntax */
        fmt.Println(strconv.ParseBool("nil"))
}
