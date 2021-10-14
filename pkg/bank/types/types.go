package types

//Shows money
type Money int

//Currency shows currency code
type Currency string

// Currency codes
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN shows card number
type PAN string

type PaymentSource struct {
	Type    string
	Number  string
	Balance Money
}

type Payment struct {
	ID     int
	Amount Money
}

//Car shows the card issue infomations
type Card struct {
	ID         int
	PAN        string
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money
}
