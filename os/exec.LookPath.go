package main

import "os/exec"
import "fmt"

func main() {
        path, error := exec.LookPath("cat")
        if error != nil {
                panic(error)
        }
        fmt.Println(path)
}
