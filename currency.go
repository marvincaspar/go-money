package money

import(
	"strconv"
	"strings"
)

// Currency represents the currency information for formatting
type Currency struct {
	code     string
	decimalDelimiter  string
	thousandDelimiter string
	exponent int
	symbol   string
	template string
}

// https://en.wikipedia.org/wiki/ISO_4217

// USD creates and returns a new Currenty instance for USD
func USD() *Currency {
	return &Currency{code: "USD", decimalDelimiter: ".", thousandDelimiter: ",", exponent: 2, symbol: "$", template: "$1"}
}

// EUR creates and returns a new Currenty instance for EUR
func EUR() *Currency {
	return &Currency{code: "EUR", decimalDelimiter: ",", thousandDelimiter: ".", exponent: 2, symbol: "€", template: "$1"}
}

// Add creates and returns a new Currenty instance
func Add(code string, decimalDelimiter string, thousandDelimiter string, exponent int, symbol string, template string) *Currency {
	return &Currency{
		code: code, 
		decimalDelimiter: decimalDelimiter, 
		thousandDelimiter: thousandDelimiter, 
		exponent: exponent, 
		symbol: symbol, 
		template: template,
	}
}

// Formatter returns currency formatter representing
// used currency structure
func (c *Currency) Format(amount int64) string {
	positiveAmount := amount
	if(amount < 0) {
		positiveAmount = amount * -1
	}
	result := strconv.FormatInt(positiveAmount, 10)

	if len(result) <= c.exponent {
		result = strings.Repeat("0", c.exponent-len(result)+1) + result
	}

	if c.thousandDelimiter != "" {
		for i := len(result) - c.exponent - 3; i > 0; i -= 3 {
			result = result[:i] + c.thousandDelimiter + result[i:]
		}
	}

	if c.exponent > 0 {
		result = result[:len(result)-c.exponent] + c.decimalDelimiter + result[len(result)-c.exponent:]
	}
	result = strings.Replace(c.template, "1", result, 1)
	result = strings.Replace(result, "$", c.symbol, 1)

	// Add minus sign for negative amount
	if amount < 0 {
		result = "-" + result
	}

	return result
}

func (c *Currency) equals(currency *Currency) bool {
	return c.code == currency.code
}
