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
