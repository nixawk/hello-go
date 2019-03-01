package main

import "time"
import "fmt"

func main() {
        date := time.Date(2019, 03, 01, 07, 36, 19, 967729223, time.UTC)
        diff := time.Now().Sub(date)  /* output: -7h48m43.67978616s */
        fmt.Println(diff)
}
