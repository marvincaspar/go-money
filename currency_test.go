package money

import "testing"

func TestCurrency_Create(t *testing.T) {
	tcs := []struct {
		currency *Currency
		code     string
	}{
		{USD(), "USD"},
		{EUR(), "EUR"},
	}

	for _, tc := range tcs {
		if tc.currency.code != tc.code {
			t.Errorf("Expected %s got %s", tc.code, tc.currency.code)
		}
	}
}
