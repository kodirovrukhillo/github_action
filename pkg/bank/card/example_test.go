package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSource() {
	cards := []types.Card{
		{
			PAN:     "5058 xxxx xxxx 0001",
			Balance: -10_000_00,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 10_000_00,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 0001",
			Balance: 10_000_00,
			Active:  false,
		},
	}
	result := PaymentSources(cards)
	fmt.Println(result[0].Number, result[0].Balance)
	// Output:
	// 5058 xxxx xxxx 8888 1000000

}
