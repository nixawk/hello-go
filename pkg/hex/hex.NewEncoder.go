package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	d := []byte{0x01, 0xff, 0x3a, 0xcd}
	writer := &bytes.Buffer{}
	hexWriter := hex.NewEncoder(writer)

	_, err := hexWriter.Write(d)
	if err != nil {
		log.Fatalf("hexWriter.Write() failed with '%s'\n", err)
	}

	encoded := writer.Bytes()
	fmt.Printf("Hex: %s\n", string(encoded))
}

/* https://www.programming-books.io/essential/go/b6e9fbb3165c4bcb907e469d86783aab-hex-base64-encoding#1 */
