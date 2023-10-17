package main

import (
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	card:=types.Card{
		ID: 12345,
		PAN: "5058 xxxx xxxx 9999",
		Balance: 9999_99,
		Currency: types.TJS,
		Active: true,
		Color: "white",
		Name: "Infinity",
	}
	fmt.Printf("%+v",card)
}