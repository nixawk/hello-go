package main

import "fmt"
import "time"

func send(c chan<- string) {
	for {
		c <- "ping"
	}
}

func recv(c <-chan string) {
	for {
		msg := <-c
		fmt.Printf("[recv]: %s\n", msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var s string

	c := make(chan string)
	go send(c)
	go recv(c)

	fmt.Scanln(&s)
}

/* Channels are the pipes that connect concurrent goroutines.
 * You can send values into channels from one goroutine and receive those
 * values into another goroutine.
 *
 * Channel Direction:
 *      Specify a direction on a channel type thus restricting to it either
*       sending or receiving.
 */
