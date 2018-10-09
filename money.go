package money

import "errors"

// Money stores amount and currency value
type Money struct {
	amount   *Amount
	currency *Currency
}

var calc = &Calculator{}

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

// Equals checkes equality between two Money instances
func (m *Money) Equals(money *Money) (bool, error) {
	if err := m.assertSameCurrency(money); err != nil {
		return false, err
	}

	return m.compare(money) == 0, nil
}

// SameCurrency checks if the given Money is equal by currency
func (m *Money) SameCurrency(money *Money) bool {
	return m.currency.equals(money.currency)
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

// IsZero returns boolean of whether the value of Money is equals to zero
func (m *Money) IsZero() bool {
	return m.Amount().val == 0
}

// IsPositive returns boolean of whether the value of Money is positive
func (m *Money) IsPositive() bool {
	return m.Amount().val > 0
}

// IsNegative returns boolean of whether the value of Money is negative
func (m *Money) IsNegative() bool {
	return m.Amount().val < 0
}

func (m *Money) assertSameCurrency(money *Money) error {
	if !m.SameCurrency(money) {
		return errors.New("Currency don't match")
	}

	return nil
}

func (m *Money) compare(money *Money) int {
	if m.Amount().val > money.Amount().val {
		return GreaterThanCheckResult
	}

	if m.Amount().val < money.Amount().val {
		return LessThanCheckResult
	}

	return EqualCheckResult
}
