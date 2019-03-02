package main

import "fmt"
import "net"
import "time"
import "encoding/gob"

func serverHandler(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("[R]", msg)
	}

	defer c.Close()
}

func server() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("listening 0.0.0.0:8888")
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go serverHandler(c)
	}
}

func client() {
	c, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := "Hello World"
	fmt.Println("[S]", msg)

	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()
}

func main() {
	go server()
	time.Sleep(time.Second * 2)
	go client()

	var s string
	fmt.Scanln(&s)
}

/*

$ $ go run net.Listen.go 
listening 0.0.0.0:8888
[S] Hello World
[R] Hello World

*/
