package goprices

import (
	"golang.org/x/text/currency"
)

// checkCurrency check if given currencyCode is valid or not
// by looking up the currencyCode in a predefined table
// if it does exist, returns string and nil error
// else return empty string and not-nil error
func checkCurrency(currencyCode string) (string, error) {
	unit, err := currency.ParseISO(currencyCode)
	if err != nil {
		return "", err
	}
	return unit.String(), nil
}

// sameKind checks if other's currency is identical to m's currency
func (m *Money) sameKind(other *Money) error {
	if m.Currency != other.Currency {
		return ErrNotSameCurrency
	}
	return nil
}
