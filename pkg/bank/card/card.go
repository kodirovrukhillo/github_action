package card

import (
	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var paymentSources []types.PaymentSource
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			paymentSources = append(paymentSources, types.PaymentSource{
				Type: "card", Number: card.PAN, Balance: card.Balance})
		}
	}
	return paymentSources
}

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{

		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: types.TJS,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}
