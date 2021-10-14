package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	var cards = []types.Card{

		{
			PAN:     "5058 xxxx xxxx 0000",
			Balance: 10_000,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 1111",
			Balance: 20_000,
			Active:  false,
		},
		{
			PAN:     "5058 xxxx xxxx 9999",
			Balance: -10_000,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 10_000,
			Active:  true,
		},
	}

	result := card.PaymentSources(cards)
	fmt.Println(result[0].Number, result[0].Balance)

}
