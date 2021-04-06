package goprices

import (
	"golang.org/x/text/currency"
)

func checkCurrency(currency_code string) (string, error) {
	unit, err := currency.ParseISO(currency_code)
	if err != nil {
		return "", err
	}
	return unit.String(), nil
}
