package main

import "fmt"


func main() {
        fmt.Println("Hello, World !");
}

/* go run hello.go */
/* go build hello.go && ./hello */

/*
 * The first line of the program package main defines the package name in which
 * this program should lie. It is a mandatory statement, as Go programs run in
 * packages. The main package is the starting point to run the program. Each
 * package has a path and name associated with it.

    debug@ulab:~/.golang$ ll go/src/fmt/
    total 204
    -rw-r--r-- 1 debug debug 14559 12月 15 07:36 doc.go
    -rw-r--r-- 1 debug debug   551 12月 15 07:36 example_test.go
    -rw-r--r-- 1 debug debug   219 12月 15 07:36 export_test.go
    -rw-r--r-- 1 debug debug 56275 12月 15 07:36 fmt_test.go
    -rw-r--r-- 1 debug debug 12673 12月 15 07:36 format.go
    -rw-r--r-- 1 debug debug 30006 12月 15 07:36 print.go
    -rw-r--r-- 1 debug debug 31951 12月 15 07:36 scan.go
    -rw-r--r-- 1 debug debug 38408 12月 15 07:36 scan_test.go
    -rw-r--r-- 1 debug debug  2156 12月 15 07:36 stringer_test.go

 */

/* https://gobyexample.com/hello-world */