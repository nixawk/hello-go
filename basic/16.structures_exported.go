package main

import "fmt"

/*
 * Struct fields whose names begin with an uppercase letter are exported.
 * All other names are unexported.
 *
 * Unexported fields can only be accessed by code within the same package.
 * As such, if you are ever accessing a field from a different package, its
 * name needs to start with an uppercase letter.
 */
type Account struct {
	UserID      int    /* exported */
	accessToken string /* unexported */
}

func main() {
	account := Account{UserID: 1, accessToken: "aZWQQQWsasqweqw"}
	fmt.Printf("%#v\n", account)
}
