package money

import "testing"

func TestMoney_New(t *testing.T) {
	m := New(1, EUR())

	if m.amount.val != 1 {
		t.Errorf("Expected %d got %d", 1, m.amount.val)
	}

	if m.currency.code != "EUR" {
		t.Errorf("Expected currency %s got %s", "EUR", m.currency.code)
	}

	m = New(-1, USD())

	if m.amount.val != -1 {
		t.Errorf("Expected %d got %d", 1, m.amount.val)
	}
}

func TestMoney_Amount(t *testing.T) {
	m := New(100, USD())

	if m.Amount().Value() != 100 {
		t.Errorf("Expected %d got %d", 1, m.Amount())
	}

	m = New(-100, USD())

	if m.Amount().Value() != -100 {
		t.Errorf("Expected %d got %d", 1, m.Amount())
	}
}

func TestMoney_Currency(t *testing.T) {
	m := New(1, EUR())

	if m.Currency().code != "EUR" {
		t.Errorf("Expected currency %s got %s", "EUR", m.Currency().code)
	}
}

func TestMoney_Equals(t *testing.T) {
	m := New(0, EUR())
	tcs := []struct {
		amount   int64
		expected bool
	}{
		{-1, false},
		{0, true},
		{1, false},
	}

	for _, tc := range tcs {
		om := New(tc.amount, EUR())
		r, err := m.Equals(om)

		if err != nil || r != tc.expected {
			t.Errorf("Expected %d Equals %d == %t got %t", m.amount.val,
				om.amount.val, tc.expected, r)
		}
	}
}

func TestMoney_SameCurrency(t *testing.T) {
	m := New(0, EUR())
	tcs := []struct {
		money    *Money
		expected bool
	}{
		{New(0, EUR()), true},
		{New(0, USD()), false},
	}

	for _, tc := range tcs {
		r := m.SameCurrency(tc.money)

		if r != tc.expected {
			t.Errorf("Expected %s same currency %s, expect %t got %t", m.Currency().code,
				tc.money.Currency().code, tc.expected, r)
		}
	}
}

func TestMoney_GreaterThan(t *testing.T) {
	m := New(10, EUR())
	tcs := []struct {
		money    *Money
		expected bool
	}{
		{New(0, EUR()), true},
		{New(100, EUR()), false},
	}

	for _, tc := range tcs {
		r, _ := m.GreaterThan(tc.money)

		if r != tc.expected {
			t.Errorf("Expected %d to be greater than %d, expect %t got %t", m.Amount().Value(),
				tc.money.Amount().Value(), tc.expected, r)
		}
	}
}

func TestMoney_GreaterThanOrEqual(t *testing.T) {
	m := New(10, EUR())
	tcs := []struct {
		money    *Money
		expected bool
	}{
		{New(0, EUR()), true},
		{New(10, EUR()), true},
		{New(100, EUR()), false},
	}

	for _, tc := range tcs {
		r, _ := m.GreaterThanOrEqual(tc.money)

		if r != tc.expected {
			t.Errorf("Expected %d to be greater than or equal %d, expect %t got %t", m.Amount().Value(),
				tc.money.Amount().Value(), tc.expected, r)
		}
	}
}

func TestMoney_LessThan(t *testing.T) {
	m := New(10, EUR())
	tcs := []struct {
		money    *Money
		expected bool
	}{
		{New(0, EUR()), false},
		{New(100, EUR()), true},
	}

	for _, tc := range tcs {
		r, _ := m.LessThan(tc.money)

		if r != tc.expected {
			t.Errorf("Expected %d to be less than %d, expect %t got %t", m.Amount().Value(),
				tc.money.Amount().Value(), tc.expected, r)
		}
	}
}

func TestMoney_LessThanOrEqual(t *testing.T) {
	m := New(10, EUR())
	tcs := []struct {
		money    *Money
		expected bool
	}{
		{New(0, EUR()), false},
		{New(10, EUR()), true},
		{New(100, EUR()), true},
	}

	for _, tc := range tcs {
		r, _ := m.LessThanOrEqual(tc.money)

		if r != tc.expected {
			t.Errorf("Expected %d to be less than or equal %d, expect %t got %t", m.Amount().Value(),
				tc.money.Amount().Value(), tc.expected, r)
		}
	}
}

func TestMoney_IsZero(t *testing.T) {
	tcs := []struct {
		money    *Money
		expected bool
	}{
		{New(-10, EUR()), false},
		{New(0, EUR()), true},
		{New(10, EUR()), false},
	}

	for _, tc := range tcs {
		r := tc.money.IsZero()

		if r != tc.expected {
			t.Errorf("Expected %d to zero, expect %t got %t", tc.money.Amount().Value(), tc.expected, r)
		}
	}
}

func TestMoney_IsPositive(t *testing.T) {
	tcs := []struct {
		money    *Money
		expected bool
	}{
		{New(-10, EUR()), false},
		{New(0, EUR()), false},
		{New(10, EUR()), true},
	}

	for _, tc := range tcs {
		r := tc.money.IsPositive()

		if r != tc.expected {
			t.Errorf("Expected %d to be positive, expect %t got %t", tc.money.Amount().Value(), tc.expected, r)
		}
	}
}

func TestMoney_IsNegative(t *testing.T) {
	tcs := []struct {
		money    *Money
		expected bool
	}{
		{New(-10, EUR()), true},
		{New(0, EUR()), false},
		{New(10, EUR()), false},
	}

	for _, tc := range tcs {
		r := tc.money.IsNegative()

		if r != tc.expected {
			t.Errorf("Expected %d to be negative, expect %t got %t", tc.money.Amount().Value(), tc.expected, r)
		}
	}
}
