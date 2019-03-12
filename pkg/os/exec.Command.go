package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func basic_command_execution() {
	cmd := exec.Command("go", "version")
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("cmd.CombinedOutput() failed with '%s'\n", err)
	}
	fmt.Printf("Output:\n%s\n", string(out))
}

func advanced_command_execution() {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", "version")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Wait() failed with '%s'\n", err)
	}

	out := append(stdout.Bytes(), stderr.Bytes()...)
	fmt.Printf("Output:\n%s\n", string(out))
}

func main() {
	basic_command_execution()
	advanced_command_execution()
}

/* https://www.programming-books.io/essential/go/c290f0566c80467a9005ab3a4024ec1d-executing-commands#1 */
