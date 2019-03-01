package main

import "strconv"
import "fmt"

func main() {
	k, _ := strconv.Atoi("1234")
	fmt.Println(k)

	_, e := strconv.Atoi("what")
	fmt.Println(e) /* strconv.Atoi: parsing "what": invalid syntax */
}
