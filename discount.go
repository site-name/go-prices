package goprices

import "github.com/site-name/decimal"

// FixedDiscount applys a fixed discount to any price type.
//
// `base` must be either *Money, *MoneyRange, *TaxedMoney, *TaxedMoneyRange
//
// Returned interface{} can be either *Money, *MoneyRange, *TaxedMoney, *TaxedMoneyRange
func FixedDiscount(base interface{}, discount *Money) (interface{}, error) {
	switch value := base.(type) {
	case *MoneyRange:
		start, err := FixedDiscount(value.Start, discount)
		if err != nil {
			return nil, err
		}
		stop, err := FixedDiscount(value.Stop, discount)
		if err != nil {
			return nil, err
		}
		return NewMoneyRange(start.(*Money), stop.(*Money))
	case *TaxedMoneyRange:
		start, err := FixedDiscount(value.Start, discount)
		if err != nil {
			return nil, err
		}
		stop, err := FixedDiscount(value.Stop, discount)
		if err != nil {
			return nil, err
		}
		return NewTaxedMoneyRange(start.(*TaxedMoney), stop.(*TaxedMoney))
	case *TaxedMoney:
		net, err := FixedDiscount(value.Net, discount)
		if err != nil {
			return nil, err
		}
		gross, err := FixedDiscount(value.Gross, discount)
		if err != nil {
			return nil, err
		}
		return NewTaxedMoney(net.(*Money), gross.(*Money))
	case *Money:
		baseSubDiscount, err := value.Sub(discount)
		if err != nil {
			return nil, err
		}
		if baseSubDiscount.Amount.GreaterThan(decimal.Zero) {
			return baseSubDiscount, nil
		}
		return NewMoney(&decimal.Zero, value.Currency)
	default:
		return nil, ErrUnknownType
	}
}

// FractionalDiscount Apply a fractional discount based on either gross or net amount
//
// `base` must be either *Money, *MoneyRange, *TaxedMoney, *TaxedMoneyRange
//
// Returned interface{} can be either *Money, *MoneyRange, *TaxedMoney, *TaxedMoneyRange
func FractionalDiscount(base interface{}, fraction *decimal.Decimal, fromGross bool) (interface{}, error) {
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
		var mul *Money
		var err error
		if fromGross {
			mul, err = value.Gross.Mul(fraction)
		} else {
			mul, err = value.Net.Mul(fraction)
		}
		if err != nil {
			return nil, err
		}
		discount, err := mul.Quantize()
		if err != nil {
			return nil, err
		}
		return FixedDiscount(value, discount)
	case *Money:
		mul, err := value.Mul(fraction)
		if err != nil {
			return nil, err
		}
		discount, err := mul.Quantize()
		if err != nil {
			return nil, err
		}
		return FixedDiscount(value, discount)

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
// Returned interface{} can be either *Money, *MoneyRange, *TaxedMoney, *TaxedMoneyRange
func PercentageDiscount(base interface{}, percentage interface{}, fromGross bool) (interface{}, error) {
	var d *decimal.Decimal
	if inter, ok := percentage.(int); ok {
		d = NewDecimal(decimal.NewFromInt32(int32(inter)))
	} else if deci, ok := percentage.(*decimal.Decimal); ok {
		d = deci
	} else {
		return nil, ErrUnknownType
	}

	factor := d.Div(decimal.NewFromInt32(100))
	return FractionalDiscount(base, &factor, fromGross)
}
