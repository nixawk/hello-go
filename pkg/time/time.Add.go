package main

import "time"
import "fmt"

func main() {
        date := time.Date(2019, 03, 01, 07, 36, 19, 967729223, time.UTC)
        diff := date.Add(8 * time.Minute)  /* output: 2019-03-01 07:44:19.967729223 +0000 UTC */
        fmt.Println(diff)
}
