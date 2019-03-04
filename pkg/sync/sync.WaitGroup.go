// This sample program demonstrates how to create goroutines and how the
// scheduler behaves.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Allocate 1 logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)

	// wg is used to wait for the program to finish.
	// Add a count of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			for c := 'a'; c <= 'z'; c++ {
				fmt.Printf("%c ", c)
			}
			fmt.Printf("\n")
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			for c := 'A'; c <= 'Z'; c++ {
				fmt.Printf("%c ", c)
			}
			fmt.Printf("\n")
		}
	}()

	// Wait for the goroutines to finish
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}
