package main

import "fmt"

type Account struct {
	NickName string
}

func (a Account) PrintNickNameValue() {
	fmt.Printf("PrintNickNameValue: %p\n", &a.NickName)
}

func (a *Account) PrintNickNamePointer() {
	fmt.Printf("PrintNickNamePointer: %p\n", &a.NickName)
}

func main() {
	account := Account{NickName: "Google"}
	fmt.Printf("Account.NickName: %p\n", &account.NickName)
	account.PrintNickNameValue()
	account.PrintNickNamePointer()
}

/* https://www.programming-books.io/essential/go/503ddd46a4854315a88973b8db780fba-methods#1 */

/*

$ go run 16.structures_value_VS_pointer.go
Account.NickName: 0xc00000e1e0
PrintNickNameValue: 0xc00000e1f0
PrintNickNamePointer: 0xc00000e1e0

*/
