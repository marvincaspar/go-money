package money

// Currency represents the currency information for formatting
type Currency struct {
	code     string
	decimal  string
	thousand string
	exponent int
	symbol   string
	template string
}

// https://en.wikipedia.org/wiki/ISO_4217

// USD creates and returns a new Currenty instance for USD
func USD() *Currency {
	return &Currency{code: "USD", decimal: ".", thousand: ",", exponent: 2, symbol: "$", template: "$1"}
}

// EUR creates and returns a new Currenty instance for EUR
func EUR() *Currency {
	return &Currency{code: "EUR", decimal: ",", thousand: ".", exponent: 2, symbol: "€", template: "$1"}
}
