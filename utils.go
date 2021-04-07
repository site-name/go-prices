package goprices

import (
	"golang.org/x/text/currency"
)

// checkCurrency check if given currency_code is valid or not
// by looking up the currency_code in a predefined table
// if it does exist, returns string and nil error
// else return empty string and not-nil error
func checkCurrency(currency_code string) (string, error) {
	unit, err := currency.ParseISO(currency_code)
	if err != nil {
		return "", err
	}
	return unit.String(), nil
}
