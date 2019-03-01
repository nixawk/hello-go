package main

import "io/ioutil"
import "fmt"

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func main() {
        data, err := ioutil.ReadFile("/tmp/testfile")
        check(err)

        fmt.Print(string(data))
}
