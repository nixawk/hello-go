package main

import "fmt"
import "io/ioutil"
import "net/http"

func main() {
        resp, err := http.Get("http://localhost:8000")
        if err != nil {
        }

        /* A defer statement defers the execution of a function until the the
           surrounding function returns.
         */
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
        }

        fmt.Print(string(body))
}

/* https://golang.org/pkg/net/http */
