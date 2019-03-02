package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) get_info() {
	fmt.Printf("name: %s, age: %d\n", p.name, p.age)
}

func main() {
	var kee Person                         /* struct without init */
	bob := Person{name: "Bob Ti", age: 25} /* struct initialization */

	kee.name = "Kee Mitz"
	kee.age = 25

	kee.get_info()
	bob.get_info()

}

/* https://gobyexample.com/structs */
/* https://tour.golang.org/moretypes/4 */
