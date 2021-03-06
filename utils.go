package goprices

import (
	"strings"

	"golang.org/x/text/currency"
)

// checkCurrency checks if given `currencyCode` is valid or not by:
//
// Checking if the money is valid by parsing
func checkCurrency(currencyCode string) (string, error) {
	unit, err := currency.ParseISO(currencyCode)
	if err != nil {
		return "", err
	}
	return unit.String(), nil
}

// SameKind checks if other's `Currency` is identical to m's `Currency`
//
// Returned error could be either `nil` or `ErrNotSameCurrency`
func (m *Money) SameKind(other *Money) error {
	if !strings.EqualFold(m.Currency, other.Currency) {
		return ErrNotSameCurrency
	}
	return nil
}

// GetCurrencyPrecision returns a number for money rounding.
//
// Returned error could be `nil` or `ErrUnknownCurrency`
//
// E.g:
//  GetCurrencyPrecision("vnd") => 0, nil
func GetCurrencyPrecision(currency string) (int32, error) {
	currencyCode, err := checkCurrency(currency)
	if err != nil {
		return 0, err
	}
	c, ok := currencies[currencyCode]
	if !ok {
		return 0, ErrUnknownCurrency
	}
	return c.Fraction, nil
}

func newInt32(in int32) *int32 {
	return &in
}
