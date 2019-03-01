package main

import "os"

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func main() {
        f, err := os.Open("/tmp/testfile")
        check(err)
        f.Close()
}
