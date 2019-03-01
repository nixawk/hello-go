package main

import "regexp"
import "bytes"
import "fmt"

func main() {
        r, _ := regexp.Compile("p[a-z]+ch")
        out :=  r.ReplaceAllFunc([]byte("peach punch"), bytes.ToUpper)
        fmt.Println(string(out))
}
