package main

import "fmt"

func init() {
	fmt.Println("app/entry.go ==> init()")
}

var myVersion = fetchVersion()

func main() {
	fmt.Println("app/fetch-version.go ==> main()")
	fmt.Println("version ===> ", myVersion)
}

/*

$ go run src/app/*.go
version/entry.go ==> getLocalVersion()
version/get-version.go ==> getVersion()
version/entry.go ==> init()
version/get-version.go ==> init()
app/fetch-version.go ==> fetchVersion()
app/entry.go ==> init()
app/fetch-version.go ==> init()
app/fetch-version.go ==> main()
version ===>  1.0.0

*/