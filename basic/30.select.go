package main

import "fmt"
import "time"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		c1 <- "one"
	}()

	go func() {
		c2 <- "two"
	}()

	for i := 0; i < 5; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(time.Second):
			fmt.Println("timeout")
		}
	}
}

/* https://gobyexample.com/select */
