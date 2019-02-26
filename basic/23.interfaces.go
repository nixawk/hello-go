package main

import "fmt"

type MyInterface interface {
        MyFunction() string
}

type MyStruct struct {
        greeting string
}

func(mystruct MyStruct) MyFunction() string {
        return mystruct.greeting
}

func PrintGreeting(myinteface MyInterface) {
        fmt.Println(myinteface.MyFunction())
}

func main() {
        var ms = MyStruct{greeting: "Hello, GoLang"}
        PrintGreeting(ms)
}

/*
 * Go programming provides another data type called interfaces which represents
 * a set of method signatures.
 */
