package main

import "fmt"

/*
for [condition |  ( init; condition; increment ) | Range] {
   statement(s);
}
*/

func foo_for() {
	var i int

	for true {
		fmt.Printf("[*] enter a int: ")
		fmt.Scanf("%d", &i)
		if i > 0 {
			break
		} else {
			fmt.Println("[!] please a positive number")
		}
	}

	for ; i < 10; i++ {
		fmt.Printf("i = %d\n", i)
	}

}

func main() {
	foo_for()
}

/* https://www.tutorialspoint.com/go/go_for_loop.htm */
