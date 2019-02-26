package main

import "fmt"

type Person struct {
        name string
        age int
}

func main() {
        var kee Person                             /* struct without init */
        var bob = Person{name: "Bob Ti", age: 25}  /* struct initialization */
        var someone *Person                        /* struct pointer */

        kee.name = "Kee Mitz"
        kee.age = 25

        someone = &kee

        fmt.Printf("name: %s, age: %d\n", kee.name, kee.age)
        fmt.Printf("name: %s, age: %d\n", bob.name, bob.age)
        fmt.Printf("name: %s, age: %d\n", someone.name, someone.age)
}

/* https://gobyexample.com/structs */
/* https://tour.golang.org/moretypes/4 */