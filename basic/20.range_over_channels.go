package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for element := range queue {
		fmt.Println(element)
	}
}

/* https://gobyexample.com/range-over-channels */
