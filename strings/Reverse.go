package main

import "fmt"

func Reverse(s string) string {
        runes := []rune(s)
        for i, l := 0, len(s); i < l / 2; i++ {
                // runes[i], runes[l - i - 1] = runes[l - i - 1], runes[i]
                runes[i] ^= runes[l - i - 1]
                runes[l - i - 1] ^= runes[i]
                runes[i] ^= runes[l - i - 1]
        }
        return string(runes)
}

func main() {
        s := "1234"
        fmt.Println(s, "-> reverse ->", Reverse(s))
}

/* https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go */
