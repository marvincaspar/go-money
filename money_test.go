package money

import (
	"reflect"
	"testing"
)

func TestMoney_New(t *testing.T) {
	m := New(1, EUR())

	if m.Amount().Value() != 1 {
		t.Errorf("Expected %d got %d", 1, m.Amount().Value())
	}

	if m.Currency().code != "EUR" {
		t.Errorf("Expected currency %s got %s", "EUR", m.Currency().code)
	}

	m = New(-1, USD())

	if m.Amount().Value() != -1 {
		t.Errorf("Expected %d got %d", 1, m.Amount().Value())
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
			t.Errorf("Expected %d Equals %d == %t got %t", m.Amount().Value(),
				om.Amount().Value(), tc.expected, r)
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

func TestMoney_Add(t *testing.T) {
	tcs := []struct {
		amount1  int64
		amount2  int64
		expected int64
	}{
		{5, 5, 10},
		{10, 5, 15},
		{1, -1, 0},
	}

	for _, tc := range tcs {
		m := New(tc.amount1, EUR())
		om := New(tc.amount2, EUR())
		r, err := m.Add(om)

		if err != nil {
			t.Error(err)
		}

		if r.Amount().Value() != tc.expected {
			t.Errorf("Expected %d + %d = %d got %d", tc.amount1, tc.amount2,
				tc.expected, r.Amount().Value())
		}
	}

}

func TestMoney_Add2(t *testing.T) {
	m := New(100, EUR())
	dm := New(100, USD())
	r, err := m.Add(dm)

	if r != nil || err == nil {
		t.Error("Expected err")
	}
}

func TestMoney_Subtract(t *testing.T) {
	tcs := []struct {
		amount1  int64
		amount2  int64
		expected int64
	}{
		{5, 5, 0},
		{10, 5, 5},
		{1, -1, 2},
	}

	for _, tc := range tcs {
		m := New(tc.amount1, EUR())
		om := New(tc.amount2, EUR())
		r, err := m.Subtract(om)

		if err != nil {
			t.Error(err)
		}

		if r.Amount().Value() != tc.expected {
			t.Errorf("Expected %d - %d = %d got %d", tc.amount1, tc.amount2,
				tc.expected, r.Amount().Value())
		}
	}
}

func TestMoney_Subtract2(t *testing.T) {
	m := New(100, EUR())
	dm := New(100, USD())
	r, err := m.Subtract(dm)

	if r != nil || err == nil {
		t.Error("Expected err")
	}
}

func TestMoney_Multiply(t *testing.T) {
	tcs := []struct {
		amount     int64
		multiplier int64
		expected   int64
	}{
		{5, 5, 25},
		{10, 5, 50},
		{1, -1, -1},
		{1, 0, 0},
	}

	for _, tc := range tcs {
		m := New(tc.amount, EUR())
		r := m.Multiply(tc.multiplier).Amount().Value()

		if r != tc.expected {
			t.Errorf("Expected %d * %d = %d got %d", tc.amount, tc.multiplier, tc.expected, r)
		}
	}
}

func TestMoney_Allocate(t *testing.T) {
	tcs := []struct {
		amount   int64
		ratios   []int
		expected []int64
	}{
		{100, []int{50, 50}, []int64{50, 50}},
		{100, []int{30, 30, 30}, []int64{34, 33, 33}},
		{200, []int{1, 1, 1}, []int64{67, 67, 66}},
		{5, []int{3, 7}, []int64{2, 3}},
	}

	for _, tc := range tcs {
		m := New(tc.amount, EUR())
		var rs []int64
		split, _ := m.Allocate(tc.ratios...)

		for _, party := range split {
			rs = append(rs, party.Amount().Value())
		}

		if !reflect.DeepEqual(tc.expected, rs) {
			t.Errorf("Expected allocation of %d for ratios %v to be %v got %v", tc.amount, tc.ratios,
				tc.expected, rs)
		}
	}
}

func TestMoney_Allocate2(t *testing.T) {
	m := New(100, EUR())
	r, err := m.Allocate()

	if r != nil || err == nil {
		t.Error("Expected err")
	}
}

func TestMoney_Comparison(t *testing.T) {
	pound := New(100, USD())
	twoPounds := New(200, USD())
	twoEuros := New(200, EUR())

	if r, err := pound.GreaterThan(twoPounds); err != nil || r {
		t.Errorf("Expected %d Greater Than %d == %t got %t", pound.Amount().Value(),
			twoPounds.Amount().Value(), false, r)
	}

	if r, err := pound.LessThan(twoPounds); err != nil || !r {
		t.Errorf("Expected %d Less Than %d == %t got %t", pound.Amount().Value(),
			twoPounds.Amount().Value(), true, r)
	}

	if r, err := pound.LessThan(twoEuros); err == nil || r {
		t.Error("Expected err")
	}

	if r, err := pound.GreaterThan(twoEuros); err == nil || r {
		t.Error("Expected err")
	}

	if r, err := pound.Equals(twoEuros); err == nil || r {
		t.Error("Expected err")
	}

	if r, err := pound.LessThanOrEqual(twoEuros); err == nil || r {
		t.Error("Expected err")
	}

	if r, err := pound.GreaterThanOrEqual(twoEuros); err == nil || r {
		t.Error("Expected err")
	}
}
