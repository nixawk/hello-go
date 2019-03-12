package main

import "plugin"

func main() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}

	h, err := p.Lookup("Hello")
	if err != nil {
		panic(err)
	}

	h.(func())()
}

/* https://www.programming-books.io/essential/go/40da6aa65709423fb43c3a8f30d24735-defining-and-using-a-plugin */
