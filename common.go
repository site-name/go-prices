package goprices

import (
	"fmt"

	"github.com/site-name/decimal"
)

type MoneyObject interface {
	*Money |
		*MoneyRange |
		*TaxedMoney |
		*TaxedMoneyRange
}

// MoneyInterface
type MoneyInterface[T MoneyObject] interface {
	Quantize(round Rounding, exp int) (T, error) // NOTE: if exp < 0, system wil use default
	fmt.Stringer
	fixedDiscount(discount Money) (T, error)
	fractionalDiscount(fraction decimal.Decimal, fromGross bool) (T, error)
	GetCurrency() string
}

// QuantizePrice accepts the `price` argument to be either:
// *TaxedMoney or *Money or *TaxedMoneyRange
func QuantizePrice[K MoneyObject, T MoneyInterface[K]](price T, round Rounding) (K, error) {
	return price.Quantize(round, -1)
}
