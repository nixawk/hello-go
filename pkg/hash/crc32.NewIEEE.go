package main

import "os"
import "fmt"
import "hash/crc32"

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("[*] Usage: %s <string>\n", os.Args[0])
		os.Exit(1)
	}

	h := crc32.NewIEEE()
	h.Write([]byte(os.Args[1]))
	v := h.Sum32()
	fmt.Println(v)
}
