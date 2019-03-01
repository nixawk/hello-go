package main

import "time"
import "fmt"

func main() {
        date := time.Date(2019, 03, 01, 07, 36, 19, 967729223, time.UTC)
        fmt.Println(date)
        fmt.Println(date.Year())
        fmt.Println(date.Month())
        fmt.Println(date.Day())
        fmt.Println(date.Hour())
        fmt.Println(date.Minute())
        fmt.Println(date.Second())
        fmt.Println(date.Nanosecond())
        fmt.Println(date.Location())

        fmt.Println(date.Weekday())

        fmt.Println(date.Before(time.Now()))
        fmt.Println(date.After(time.Now()))
        fmt.Println(date.Equal(time.Now()))
}