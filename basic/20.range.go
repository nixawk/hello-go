package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}                      /* arrary */
	b := map[string]string{"k1": "v1", "k2": "v2"} /* map */

	for i := range a {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}

	for k := range b {
		fmt.Printf("b[\"%s\"] = \"%s\"\n", k, b[k])
	}
}

/*
 * The range keyword is used in for loop to iterate over items of an array,
 * slice, channel, or map.
 * With array and slices, it returns the index of the item as integer.
 * With maps, it returns the key of the next key-value pair.
 * Range either returns one value or two.
 */

/* https://www.tutorialspoint.com/go/go_range.htm */
