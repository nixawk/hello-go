package main

import (
        "io"
        "fmt"
        "crypto/sha1"
)

func main() {
        h := sha1.New()
        io.WriteString(h, "admin")
        fmt.Printf("%x\n", h.Sum(nil))  /* d033e22ae348aeb5660fc2140aec35850c4da997 */
}

/* https://golang.org/pkg/crypto/sha1/ */
