package main

import "fmt"

func main() {
        messages := make(chan string)       /* create a new channel with make */
        go func() { messages <- "ping" }()  /* send a value into a channel */

        msg := <-messages                   /* receive a value from the channel */
        fmt.Println(msg)
}

/* Channels are the pipes that connect concurrent goroutines. 
 * You can send values into channels from one goroutine and receive those
 * values into another goroutine.
 */