package main

import (
	"flag"
	"fmt"
	"os"
)

var lang string /* commandline option */

func init() {

	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Custom Usage of %s:\n", os.Args[0])
	// 	flag.PrintDefaults()
	// }

	flag.CommandLine = flag.NewFlagSet("question", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Custom Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&lang, "lang", "go", "golang")
}

func main() {
	flag.Parse()
	fmt.Println("This is a console prog for", lang)
}
