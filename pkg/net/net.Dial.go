package main

import (
	"fmt"
	"net"
)

func SendTcpMsg(host string, port int, msg string) {
	c, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}
	defer c.Close()

	nwrite, err := c.Write([]byte(msg))
	if err != nil {
		panic(err)
	}

	fmt.Printf("[*] sent %d bytes\n", nwrite)
}

func RecvTcpMsg(host string, port int) {
	c, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}
	defer c.Close()

	buff := make([]byte, 1024)
	nread, err := c.Read(buff)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[*] recv %d bytes: %s\n", nread, buff[:nread])
}

func main() {
	// $ echo AAAAAA  | ncat -v -l -p -k 4444
	RecvTcpMsg("localhost", 4444)
	SendTcpMsg("localhost", 4444, "hello, golang")
}
