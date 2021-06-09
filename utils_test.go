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
		{"usD", "USD"},
		{"dkk", "DKK"},
	}

	for i, v := range values {
		unit, err := checkCurrency(v.input)
		if err != nil {
			t.Fatalf("Failed at test case: %d, error: %q", i+1, err.Error())
		}
		if unit != v.expected {
			t.Fatalf("Failed at test case: %d, got: %q, expected: %q", i+1, unit, v.expected)
		}
	}
}

func TestGetCurrencyPrecision(t *testing.T) {
	type testUnit struct {
		currency string
		expected int
	}
	testCases := []testUnit{
		{VND, 0},
		{USD, 2},
		{DKK, 2},
	}
	for index, test := range testCases {
		fraction, err := GetCurrencyPrecision(test.currency)
		if err != nil {
			t.Fatalf("Error GetCurrencyPrecision at index: %d, err: %v", index, err)
		}
		if fraction != test.expected {
			t.Fatalf("Error at index: %d, expected: %d, got: %d", index, test.expected, fraction)
		}
	}
}
