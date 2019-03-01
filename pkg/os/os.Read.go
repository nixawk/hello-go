package main

import "os"
import "fmt"

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func main() {
        data := make([]byte, 5)

        f, err := os.Open("/tmp/testfile")
        check(err)

        n, err := f.Read(data)
        check(err)

        fmt.Printf("read %d bytes. (%s)", n, data)

        f.Close()
}

/* https://github.com/golang/go/blob/master/src/os/file.go#L104 */
