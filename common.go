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
	fractionalDiscount(fraction decimal.Decimal, fromGross bool) (T, error)
	fixedDiscount(discount *Money) (T, error)
	MyCurrency() string
	Equal(T) bool
	LessThan(T) bool
	LessThanOrEqual(T) bool
	fmt.Stringer
	// Add(any) T
	// Sub(any) T
	// Mul(any) (T, error)
	// TrueDiv(any) (T, error)
}

// QuantizePrice accepts the `price` argument to be either:
// *TaxedMoney or *Money or *TaxedMoneyRange
func QuantizePrice[K MoneyObject, T MoneyInterface[K]](price T, round Rounding) (K, error) {
	return price.Quantize(round, -1)
}
