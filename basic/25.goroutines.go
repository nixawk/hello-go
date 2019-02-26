package main

import "fmt"


func f(from string) {
        for i := 0; i < 4; i++ {
                fmt.Println(from, ":", i)
        }
}

func main() {
        f("sync")

        go f("goroutine")
        go func(msg string) {
                fmt.Println(msg)
        }("going")

        fmt.Scanln()
        fmt.Println("done")
}

/* A goroutine is a lightweight thread of execution */
/*
 * Suppose we have a function call f(s). Here's how we'd call that in the usual
 * way, running it synchronously.

 * To invoke this function in a goroutine, use go f(s).
 * This new goroutine will execute concurrently with the calling one.
 */