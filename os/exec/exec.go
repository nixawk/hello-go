package main

import (
        "fmt"
        "os/exec"
)

func main() {
        output := exec.Command("/bin/bash", "-c", "pwd")
        fmt.Print(output)
}