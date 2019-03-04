package main

import "encoding/json"
import "fmt"

type response1 struct {
	Page   int
	Fruits []string
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB)) /* output: true */

	intB, _ := json.Marshal(123)
	fmt.Println(string(intB)) /* output: 123 */

	fltB, _ := json.Marshal(1.23)
	fmt.Println(string(fltB)) /* output: 1.23 */

	strB, _ := json.Marshal("golang")
	fmt.Println(string(strB)) /* output: "golang" */

	slcD := []string{"XX", "YY", "ZZ"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB)) /* output: ["XX","YY","ZZ"] */

	mapD := map[string]int{"XX": 11, "YY": 22, "ZZ": 33}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB)) /* output: {"XX":11,"YY":22,"ZZ":33} */

	res1D := &response1{
		Page:   1,
		Fruits: []string{"XX", "YY", "ZZ"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B)) /* output: {"Page":1,"Fruits":["XX","YY","ZZ"]} */
}

/*

$ go run json.Marshal.go
true
123
1.23
"golang"
["XX","YY","ZZ"]
{"XX":11,"YY":22,"ZZ":33}
{"Page":1,"Fruits":["XX","YY","ZZ"]}

*/
