package money

// Amount stores the amounts used for calculations
type Amount struct {
	val int64
}

// Money stores amount and currency value
type Money struct {
	amount   *Amount
	currency *Currency
}

// New creates and returns a new Money instance
func New(amount int64, currency *Currency) *Money {
	return &Money{
		amount:   &Amount{val: amount},
		currency: currency,
	}
}
