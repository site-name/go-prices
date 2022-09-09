package goprices

import "github.com/site-name/decimal"

// FixedDiscount applys a fixed discount to any price type.
//
// `base` must be either *Money, *MoneyRange, *TaxedMoney, *TaxedMoneyRange
//
// Returned interface{} can be either *Money, *MoneyRange, *TaxedMoney, *TaxedMoneyRange
//
// the returned interface's type should be identical to given base's type
func FixedDiscount[K MoneyObject, T MoneyInterface[K]](base T, discount *Money) (K, error) {
	return base.FixedDiscount(discount)
}

// FractionalDiscount Apply a fractional discount based on either gross or net amount
//
// `base` must be either *Money, *MoneyRange, *TaxedMoney, *TaxedMoneyRange
//
// Returned interface{} can be either *Money, *MoneyRange, *TaxedMoney, *TaxedMoneyRange
func FractionalDiscount(base interface{}, fraction decimal.Decimal, fromGross bool) (interface{}, error) {
	switch value := base.(type) {
	case *MoneyRange:
		start, err := FractionalDiscount(value.Start, fraction, fromGross)
		if err != nil {
			return nil, err
		}
		stop, err := FractionalDiscount(value.Stop, fraction, fromGross)
		if err != nil {
			return nil, err
		}
		return NewMoneyRange(start.(*Money), stop.(*Money))

	case *TaxedMoneyRange:
		start, err := FractionalDiscount(value.Start, fraction, fromGross)
		if err != nil {
			return nil, err
		}
		stop, err := FractionalDiscount(value.Stop, fraction, fromGross)
		if err != nil {
			return nil, err
		}
		return NewTaxedMoneyRange(start.(*TaxedMoney), stop.(*TaxedMoney))

	case *TaxedMoney:
		var (
			mul *Money
			err error
		)
		if fromGross {
			mul, err = value.Gross.Mul(fraction)
		} else {
			mul, err = value.Net.Mul(fraction)
		}
		if err != nil {
			return nil, err
		}
		discount, err := mul.Quantize(nil, Down)
		if err != nil {
			return nil, err
		}
		return FixedDiscount[*TaxedMoney](value, discount)

	case *Money:
		mul, err := value.Mul(fraction)
		if err != nil {
			return nil, err
		}
		discount, err := mul.Quantize(nil, Down)
		if err != nil {
			return nil, err
		}
		return FixedDiscount[*Money](value, discount)

	default:
		return nil, ErrUnknownType
	}
}

// PercentageDiscount Apply a percentage discount based on either gross or net amount.
//
// `base` must be either *Money, *MoneyRange, *TaxedMoney, *TaxedMoneyRange
//
// `percentage` must be either int or *Decimal
//
// Returned interface's type should be identical to base's type
func PercentageDiscount(base interface{}, percentage interface{}, fromGross bool) (interface{}, error) {
	var d decimal.Decimal

	switch t := percentage.(type) {
	case int:
		d = decimal.NewFromInt32(int32(t))
	case decimal.Decimal:
		d = t

	default:
		return nil, ErrUnknownType
	}

	factor := d.Div(decimal.NewFromInt32(100))
	return FractionalDiscount(base, factor, fromGross)
}
