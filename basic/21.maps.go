package main

import "fmt"

func main() {
        var people map[string]string

        people = make(map[string]string)

        people["name"] = "Matz Giea"
        people["flag"] = "AAzz"
        people["dog"] = "Tedon"

        for k := range people {
                fmt.Printf("%s - %s\n", k, people[k])
        }
}

/* var map_var[key_data_type]value_data_type */