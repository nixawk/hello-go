package main

import "fmt"
import "time"
import "math/rand"

func f(from string) {
	for i := 0; i < 4; i++ {
		fmt.Println(from, ":", i)
		amt := time.Duration(rand.Intn(1000))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 3; i++ {
		go f(fmt.Sprintf("gr%d", i))
	}
	var input string
	fmt.Scanln(&input)
}

/* A goroutine is a lightweight thread of execution */
/*
 * Suppose we have a function call f(s). Here's how we'd call that in the usual
 * way, running it synchronously.

 * To invoke this function in a goroutine, use go f(s).
 * This new goroutine will execute concurrently with the calling one.
 */
