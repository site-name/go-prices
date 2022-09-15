package goprices

import (
	"fmt"

	"github.com/site-name/decimal"
)

type MoneyObject interface {
	*Money | *MoneyRange | *TaxedMoney | *TaxedMoneyRange
}

// MoneyInterface
type MoneyInterface[T MoneyObject] interface {
	fmt.Stringer
	Quantize(*int, Rounding) (T, error)
	MyCurrency() string
	Add(any) (T, error)
	Sub(any) (T, error)
	Mul(any) (T, error)
	TrueDiv(any) (T, error)
	Equal(T) bool
	LessThan(T) bool
	LessThanOrEqual(T) bool
	fixedDiscount(discount *Money) (T, error)
	fractionalDiscount(fraction decimal.Decimal, fromGross bool) (T, error)
}

// QuantizePrice accepts the `price` argument to be either:
// *TaxedMoney or *Money or *TaxedMoneyRange
func QuantizePrice[K MoneyObject, T MoneyInterface[K]](price T, round Rounding) (K, error) {
	return price.Quantize(nil, round)
}
