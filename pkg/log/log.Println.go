// This sample program demonstrates how to use the base log package.
package main

import "log"

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println writes to the standard logger.
	log.Println("message")

	// Fatalln is Println() followed by a call to os.Exit(1).
	log.Fatalln("fatal message")

	// Panicln is Println() followed by a call to panic().
	log.Panicln("panic message")
}

/*

$ go run /tmp/log.Println.go
TRACE: 2019/03/04 18:04:39.190193 /tmp/log.Println.go:13: message
TRACE: 2019/03/04 18:04:39.190377 /tmp/log.Println.go:16: fatal message
exit status 1

*/
