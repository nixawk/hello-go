package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	d := []byte{0x01, 0xff, 0x3a, 0xcd}
	s := hex.EncodeToString(d)
	fmt.Printf("Hex: %s\n", s)

	d2, err := hex.DecodeString(s)
	if err != nil {
		log.Fatalf("hex.DecodeString() failed with '%s'\n", err)
	}

	if !bytes.Equal(d, d2) {
		log.Fatalf("decoded version is different than original")
	}
}

/* https://www.programming-books.io/essential/go/b6e9fbb3165c4bcb907e469d86783aab-hex-base64-encoding#1 */
