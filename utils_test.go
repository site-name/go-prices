package goprices

import (
	"testing"
)

func TestCheckCurrency(t *testing.T) {
	type testUnit struct {
		input    string
		expected string
	}
	var values = []testUnit{
		{"vnd", "VND"},
		{"usd", "USD"},
		{"dkk", "DKL"}, // error
	}

	for i, v := range values {
		unit, err := checkCurrency(v.input)
		if err != nil {
			t.Fatalf("Failed at test case: %d, error: %q", i, err.Error())
		}
		if unit != v.expected || unit == "" {
			t.Fatalf("Failed at test case: %d, got: %q, expected: %q", i, unit, v.expected)
		}
	}
}
