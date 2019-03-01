package main

import "os"
import "fmt"

func main() {
        for k, v := range os.Args {
                fmt.Println(k, v)
        }
}

/*

$ go run os.Args.go a b c d e f g h i j
0 /tmp/go-build065938888/b001/exe/os.Args
1 a
2 b
3 c
4 d
5 e
6 f
7 g
8 h
9 i
10 j

*/
