package main

import "io/ioutil"

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func main() {
        data := "Hello Golang"
        err := ioutil.WriteFile("/tmp/testfile", []byte(data), 0600)
        check(err)
}
