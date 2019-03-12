package main

import "os"
import "fmt"

func check_err(e error) {
	if e != nil {
		panic(e)
	}
}

func defer_foo() {
	fmt.Println("1st")
	defer fmt.Println("4th")
	defer fmt.Println("3rd")
	fmt.Println("2nd")
}

func defer_exit_foo() {
	fmt.Println("AAA")
	os.Exit(0)
	defer fmt.Println("BBB") /* never executed*/
	fmt.Println("CCC")
}

func defer_file_foo() {
	f, err := os.Open("/etc/passwd")
	check_err(err)

	defer f.Close()

	d := make([]byte, 32)
	n, err := f.Read(d)
	check_err(err)

	fmt.Printf("\n[*] read %d bytes -> %s\n", n, string(d))
}

func main() {
	defer_foo()
	defer_file_foo()
	defer_exit_foo()
}

/* Go has a special statement called defer which shecules a function call to
   be run after the function completes.
   defer is often used when resources need to be freed in some way. */

/*

$ go run defer.go
1st
2nd
3rd
4th

[*] read 32 bytes -> root:x:0:0:root:/root:/bin/bash

AAA

*/
