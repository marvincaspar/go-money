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


func TestCurrency_Add(t *testing.T) {
	c := Add("BTC", ".", "", 8, "BTC", "1BTC")

	if c.code != "BTC" {
		t.Errorf("Expected %s got %s", "BTC", c.code)
	}
}
