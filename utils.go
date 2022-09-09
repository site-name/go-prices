package goprices

import (
	"strings"

	"golang.org/x/text/currency"
)

// validateCurrency checks if given `currencyCode` is valid or not.
// When it is not, returns empty string and ErrUnknownCurrency
func validateCurrency(currencyCode string) (string, error) {
	unit, err := currency.ParseISO(currencyCode)
	if err != nil {
		return "", ErrUnknownCurrency
	}
	return unit.String(), nil
}

// SameKind checks if other's currency is identical to current money currency.
// If other is nil, returns false.
func (m *Money) SameKind(other *Money) bool {
	if other == nil {
		return false
	}
	return strings.EqualFold(m.Currency, other.Currency)
}

// GetCurrencyPrecision returns a number for money rounding.
//
// Returned error could be `nil` or `ErrUnknownCurrency`
//
// E.g:
//
//	GetCurrencyPrecision("vnd") => 0, nil
func GetCurrencyPrecision(currency string) (int, error) {
	currencyCode, err := validateCurrency(currency)
	if err != nil {
		return 0, err
	}
	c, ok := currencies[currencyCode]
	if !ok {
		return 0, ErrUnknownCurrency
	}
	return c.Fraction, nil
}

func newInt(in int) *int {
	return &in
}
