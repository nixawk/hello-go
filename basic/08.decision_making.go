package main

import "fmt"


func foo_if() {
        var a int  /* local variable definition */

        fmt.Printf("[*] (foo_if) enter a int: ")
        fmt.Scanf("%d", &a)

        if (a < 20) {
                fmt.Printf("%d is less than 20\n", a)
        } else {
                fmt.Printf("%d is greater than 20\n", a)
        }
}


func foo_switch () {
        var b int;  /* local variable definition */

        fmt.Printf("[*] (foo_switch) enter a int: ")
        fmt.Scanf("%d", &b)

        switch (b) {
        case 1:
                fmt.Println("C programming")
                break
        case 2:
                fmt.Println("Python programming")
                break
        default:
                fmt.Println("Go programming")
                break
        }

}

func main() {
        foo_if()
        foo_switch()
}

/* if statement */
/* if ... else statement */
/* nested if statements */
/* switch statement */
/* select statement */