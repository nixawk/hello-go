package main

import "fmt"

type Painter interface {
	Paint()
}

type Rembrandt struct{}
type Picasso struct{}

func (r Rembrandt) Paint() {}
func (r Picasso) Paint()   {}

func WhichPainter(painter Painter) {
	switch painter.(type) {
	case Rembrandt:
		fmt.Println("The underlying type is Rembrandt")
	case Picasso:
		fmt.Println("The underlying type is Picasso")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	r := new(Rembrandt)
	WhichPainter(r)

	p := new(Picasso)
	WhichPainter(p)
}

/* https://www.programming-books.io/essential/go/d2b48c54f8434da2bcc3bead44ef81f6-determining-underlying-type-from-interface */
