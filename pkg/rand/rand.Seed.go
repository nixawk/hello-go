// This sample program demonstrates how to use an unbuffered channel to
// simulate a game of tennis between two goroutines.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// player simulates a person playing the game of tennis.
func player(name string, court chan int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n := rand.Intn(100)
		if n%3 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		court <- ball
	}
}

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// main is the entry point for all Go programs.
func main() {
	// Create an unbuffered channel.
	court := make(chan int)

	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Launch two players.
	go player("Nadal", court)
	go player("Djokovic", court)

	// Start the set.
	court <- 1

	// Wait for  the game to finish.
	wg.Wait()
}