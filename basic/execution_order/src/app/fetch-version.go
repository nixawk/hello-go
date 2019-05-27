package main

import (
	"fmt"
	"version"
)

func init() {
	fmt.Println("app/fetch-version.go ==> init()")
}

func fetchVersion() string {
	fmt.Println("app/fetch-version.go ==> fetchVersion()")
	return version.Version
}
