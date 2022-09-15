package goprices

import "github.com/site-name/decimal"

// FixedDiscount applys a fixed discount to any price type.
func FixedDiscount[K MoneyObject, T MoneyInterface[K]](base T, discount *Money) (K, error) {
	return base.fixedDiscount(discount)
}

// FractionalDiscount Apply a fractional discount based on either gross or net amount
func FractionalDiscount[K MoneyObject, T MoneyInterface[K]](base T, fraction decimal.Decimal, fromGross bool) (K, error) {
	return base.fractionalDiscount(fraction, fromGross)
}

// PercentageDiscount Apply a percentage discount based on either gross or net amount.
func PercentageDiscount[K MoneyObject, T MoneyInterface[K]](base T, percentage float64, fromGross bool) (K, error) {
	factor := decimal.NewFromFloat(percentage).Div(decimal.NewFromFloat(100))
	return base.fractionalDiscount(factor, fromGross)
}
