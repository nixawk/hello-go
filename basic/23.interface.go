package main

import "fmt"

type Singer interface {
	Sing()
}

type Writer interface {
	Write()
}

type Human struct{}

func (h *Human) Sing() {
	fmt.Println("singing")
}

func (h *Human) Write() {
	fmt.Println("writing")
}

type OnlySinger struct{}

func (o *OnlySinger) Sing() {
	fmt.Println("singing")
}

func main() {
	h := new(Human)
	h.Sing()
	h.Write()

	o := new(OnlySinger)
	o.Sing()
}

/* https://www.programming-books.io/essential/go/23c5afaeaad2401bb967b117a0edb5bb-simple-interface */
