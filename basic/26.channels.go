package main

import (
	"log"
	"sync"
)

// A WaitGroup waits for a collection of goroutines to finish. The main
// goroutine calls Add to set the number of goroutines to wait for . Then each
// of the goroutines runs and calls Done when finished. At the same time, Wait
// can be used to block until all goroutines have finished.
var wg sync.WaitGroup

func readchannel(ch chan int) {
	for i := range ch {
		log.Println(i)
	}

	defer wg.Done()
}

func main() {
	ch := make(chan int)

	wg.Add(1)

	go readchannel(ch)

	for i := 1; i < 10; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()
}

/* Channels are the pipes that connect concurrent goroutines.
 * You can send values into channels from one goroutine and receive those
 * values into another goroutine.
 *
 * Channel Direction:
 *      Specify a direction on a channel type thus restricting to it either
 *      sending or receiving.
 */
