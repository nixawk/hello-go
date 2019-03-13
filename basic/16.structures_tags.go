package main

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Username      string `json:"username"`
	DisplayName   string `json:"display_name"`
	FavoriteColor string `json:"favorite_color,omitempty"`
}

func main() {
	json_str := `
        {
                "username": "debug",
                "display_name": "Smith Bob"
        }`

	account := new(Account)
	json.Unmarshal([]byte(json_str), account)
	fmt.Println(account)

	new_json, _ := json.Marshal(account)
	fmt.Printf("%s\n", new_json)
}

/*
 * Struct fields can have tags associated with them. These tags can be read by
 * the reflect package to get custom information specified about a field by the
 * developer.
 */
/* https://www.programming-books.io/essential/go/fc4cf2bb9b2345a9a9fe53b4ed65ae12-struct-tags */
/* https://stackoverflow.com/questions/10858787/what-are-the-uses-for-tags-in-go */
