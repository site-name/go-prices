package goprices

import (
	"github.com/site-name/decimal"
)

type MoneyObject interface {
	Money |
		MoneyRange |
		TaxedMoney |
		TaxedMoneyRange
}

type MoneyInterface[T MoneyObject] interface {
	String() string
	GetCurrency() string
	Quantize(round Rounding, exp int) (*T, error) // NOTE: if exp < 0, system wil use default
	fixedDiscount(discount Money) (*T, error)
	fractionalDiscount(fraction decimal.Decimal, fromGross bool, rounding Rounding) (*T, error)
	Neg() T
}

// QuantizePrice accepts the `price` argument to be either:
// TaxedMoney or Money or TaxedMoneyRange
func QuantizePrice[K MoneyObject, T MoneyInterface[K]](price T, round Rounding) (*K, error) {
	return price.Quantize(round, -1)
}
