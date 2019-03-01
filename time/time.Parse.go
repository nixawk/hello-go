package main

import "time"
import "fmt"

func main() {
        now := time.Now()
        fmt.Println(now.Format(time.RFC3339)) /* output: 2019-03-01T08:01:57+08:00 */
}
