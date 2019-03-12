package main

import "fmt"

type ByteSize float64

const (
	Low = iota
	Medium
	High
)

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Printf("Low: %d\nMedium: %d\nHigh: %d\n", Low, Medium, High)
	fmt.Printf("KB: %f\nMB: %f\nGB: %f\n", KB, MB, GB)
}

/* https://github.com/golang/go/wiki/Iota */
/* https://www.programming-books.io/essential/go/58a7d48d4d59472a9a7e6fb561771f8d-iota */
