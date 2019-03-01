package main

import "flag"
import "fmt"

func main() {
        wordPtr := flag.String("word", "foo", "a string")
        numbPtr := flag.Int("numb", 42, "an int")
        boolPtr := flag.Bool("fork", false, "a bool")

        var svar string
        flag.StringVar(&svar, "svar", "bar", "a string var")
        flag.Parse()

        fmt.Println("word:", *wordPtr)
        fmt.Println("numb:", *numbPtr)
        fmt.Println("fork:", *boolPtr)
        fmt.Println("svar:", svar)
        fmt.Println("tail:", flag.Args())

}

/* https://gobyexample.com/command-line-flags */

/*

$ go run flag.String.go --word=aaa --numb=1234 --fork=true --svar=xxx
word: aaa
numb: 1234
fork: true
svar: xxx
tail: []

*/
