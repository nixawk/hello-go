package main

import "time"
import "fmt"

func main() {
        now := time.Now()
        fmt.Println(now.Format("3:00PM"))
}
