package money

import "errors"

// Money stores amount and currency value
type Money struct {
	amount   *Amount
	currency *Currency
}

var calc = &calculator{}

const (
	GreaterThanCheckResult = 1
	EqualCheckResult       = 0
	LessThanCheckResult    = -1
)

// New creates and returns a new Money instance
func New(amount int64, currency *Currency) *Money {
	return &Money{
		amount:   &Amount{val: amount},
		currency: currency,
	}
}

// Currency returns the currency used by Money
func (m *Money) Currency() *Currency {
	return m.currency
}

// Amount returns the amount value as int64
func (m *Money) Amount() *Amount {
	return m.amount
}

// Add returns new Money struct with value representing sum of Self and Other Money
func (m *Money) Add(money *Money) (*Money, error) {
	if err := m.assertSameCurrency(money); err != nil {
		return nil, err
	}

	return &Money{amount: calc.add(m.Amount(), money.Amount()), currency: m.currency}, nil
}

// Subtract returns new Money struct with value representing difference of Self and Other Money
func (m *Money) Subtract(money *Money) (*Money, error) {
	if err := m.assertSameCurrency(money); err != nil {
		return nil, err
	}

	return &Money{amount: calc.subtract(m.Amount(), money.Amount()), currency: m.currency}, nil
}

// Multiply returns new Money struct with value representing Self multiplied value by multiplier
func (m *Money) Multiply(mul int64) *Money {
	return &Money{amount: calc.multiply(m.Amount(), mul), currency: m.currency}
}

// Allocate returns slice of Money structs with split Self value in given ratios.
// It lets split money by given ratios without losing pennies and as Split operations distributes
// leftover pennies amongst the parties with round-robin principle.
func (m *Money) Allocate(rs ...int) ([]*Money, error) {
	if len(rs) == 0 {
		return nil, errors.New("No ratios specified")
	}

	// Calculate total of ratios
	var total int
	for _, r := range rs {
		total += r
	}

	var remainder = m.Amount().Value()
	var ms []*Money
	for _, r := range rs {
		m := &Money{
			amount:   calc.allocate(m.Amount(), r, total),
			currency: m.currency,
		}

		ms = append(ms, m)
		remainder -= m.Amount().Value()
	}

	for i := 0; i < int(remainder); i++ {
		ms[i] = &Money{
			amount:   calc.add(ms[i].Amount(), &Amount{1}),
			currency: ms[i].Currency(),
		}
	}

	return ms, nil
}

// Equals checkes equality between two Money instances
func (m *Money) Equals(money *Money) (bool, error) {
	if err := m.assertSameCurrency(money); err != nil {
		return false, err
	}

	return m.compare(money) == 0, nil
}

// GreaterThan checks whether the value of Money is greater than the other
func (m *Money) GreaterThan(money *Money) (bool, error) {
	if err := m.assertSameCurrency(money); err != nil {
		return false, err
	}

	return m.compare(money) == GreaterThanCheckResult, nil
}

// GreaterThanOrEqual checks whether the value of Money is greater or equal than the other
func (m *Money) GreaterThanOrEqual(money *Money) (bool, error) {
	if err := m.assertSameCurrency(money); err != nil {
		return false, err
	}

	return m.compare(money) >= EqualCheckResult, nil
}

// LessThan checks whether the value of Money is less than the other
func (m *Money) LessThan(money *Money) (bool, error) {
	if err := m.assertSameCurrency(money); err != nil {
		return false, err
	}

	return m.compare(money) == LessThanCheckResult, nil
}

// LessThanOrEqual checks whether the value of Money is less or equal than the other
func (m *Money) LessThanOrEqual(money *Money) (bool, error) {
	if err := m.assertSameCurrency(money); err != nil {
		return false, err
	}

	return m.compare(money) <= EqualCheckResult, nil
}

func (m *Money) assertSameCurrency(money *Money) error {
	if !m.currency.equals(money.currency) {
		return errors.New("Currency don't match")
	}

	return nil
}

func (m *Money) compare(money *Money) int {
	if m.Amount().Value() > money.Amount().Value() {
		return GreaterThanCheckResult
	}

	if m.Amount().Value() < money.Amount().Value() {
		return LessThanCheckResult
	}

	return EqualCheckResult
}
