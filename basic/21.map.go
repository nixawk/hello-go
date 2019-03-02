package main

import "fmt"

func pmap(m map[string]string) {
	for k, v := range m {
		fmt.Printf("%s = %s\n", k, v)
	}
	fmt.Println("--------------------------------------")
}

func main() {

	/* panic: assignment to entry in nil map */
	/* maps have to be intialized before they can be used */
	map1 := make(map[string]string)

	map1["k1"] = "v1"
	map1["k2"] = "v2"
	map1["k3"] = "v3"

	map2 := map[string]string{
		"K1": "V1",
		"K2": "V2",
		"K3": "V3"}

	/* delete */
	pmap(map1)
	delete(map1, "k3")
	pmap(map1)

	pmap(map2)
	delete(map2, "K1")
	pmap(map2)

	fmt.Println(map1["k3"] == "")   /* true */
	fmt.Println(map2["K1"] == "V1") /* false */
}

/* var map_var[key_data_type]value_data_type */
