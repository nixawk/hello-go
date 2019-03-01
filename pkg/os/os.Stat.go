package main

import "os"
import "fmt"

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func main() {
        fi, err := os.Stat("/tmp/testfile")
        check(err)


        fmt.Printf("size : %d\n", fi.Size())
}

/*

type FileInfo interface {
        Name() string       // base name of the file
        Size() int64        // length in bytes for regular files; system-dependent for others
        Mode() FileMode     // file mode bits
        ModTime() time.Time // modification time
        IsDir() bool        // abbreviation for Mode().IsDir()
        Sys() interface{}   // underlying data source (can return nil)
}

*/
/* https://golang.org/pkg/os/#File.Stat */
/* https://golang.org/pkg/os/#FileInfo */
