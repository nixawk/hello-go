package main

import "crypto/sha1"
import "fmt"

func sha1sum(s string) string {
        h := sha1.New()
        h.Write([]byte(s))
        return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
        s := "hello world"

        fmt.Println(sha1sum(s))
}

/* https://stackoverflow.com/questions/19328957/golang-from-bytes-to-get-hexadecimal */
