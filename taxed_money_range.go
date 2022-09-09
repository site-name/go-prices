package goprices

import (
	"fmt"

	"github.com/site-name/decimal"
)

type TaxedMoneyRange struct {
	Start    *TaxedMoney
	Stop     *TaxedMoney
	Currency string
}

var (
	_ Currencyable                     = (*TaxedMoneyRange)(nil)
	_ MoneyInterface[*TaxedMoneyRange] = (*TaxedMoneyRange)(nil)
)

// NewTaxedMoneyRange create new taxed money range.
// It returns nil and error value if start > stop or they have different currencies
func NewTaxedMoneyRange(start, stop *TaxedMoney) (*TaxedMoneyRange, error) {
	if start == nil || stop == nil {
		return nil, ErrNillValue
	}
	_, err := validateCurrency(start.Currency)
	if err != nil {
		return nil, err
	}
	unit, err := validateCurrency(stop.Currency)
	if err != nil {
		return nil, err
	}
	if start.Net.Amount.LessThan(decimal.Zero) || stop.Gross.Amount.LessThan(decimal.Zero) {
		return nil, ErrMoneyNegative
	}

	if stop.LessThan(start) {
		return nil, ErrStopLessThanStart
	}

	return &TaxedMoneyRange{start, stop, unit}, nil
}

// String implements fmt.Stringer interface
func (t *TaxedMoneyRange) String() string {
	return fmt.Sprintf("TaxedMoneyRange{%s, %s}", t.Start.String(), t.Stop.String())
}

// MyCurrency returns current taxed money range's Currency
func (t *TaxedMoneyRange) MyCurrency() string {
	return t.Currency
}

// Add adds this taxed money range to another value
// other must be either: *Money, *MoneyRange or *TaxedMoneyRange or *TaxedMoney
func (t *TaxedMoneyRange) Add(other interface{}) (*TaxedMoneyRange, error) {
	switch v := other.(type) {
	case *Money, *TaxedMoney:
		start, err := t.Start.Add(v)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Add(v)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{start, stop, t.Currency}, nil
	case *MoneyRange:
		start, err := t.Start.Add(v.Start)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Add(v.Stop)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{start, stop, t.Currency}, nil
	case *TaxedMoneyRange:
		start, err := t.Start.Add(v.Start)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Add(v.Stop)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{start, stop, t.Currency}, nil
	default:
		return nil, ErrUnknownType
	}
}

// Sub substract this taxed money range to given other.
// other must be either *Money or *TaxedMoney or *MoneyRange or *TaxedMoneyRange
func (t *TaxedMoneyRange) Sub(other interface{}) (*TaxedMoneyRange, error) {
	if other == nil {
		return nil, ErrNillValue
	}

	switch v := other.(type) {
	case *Money, *TaxedMoney:
		start, err := t.Start.Sub(v)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Sub(v)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{start, stop, t.Currency}, nil

	case *MoneyRange:
		start, err := t.Start.Sub(v.Start)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Sub(v.Stop)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{start, stop, t.Currency}, nil

	case *TaxedMoneyRange:
		start, err := t.Start.Sub(v.Start)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Sub(v.Stop)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{start, stop, t.Currency}, nil

	default:
		return nil, ErrUnknownType
	}
}

// Equal compares two taxed money range
func (t *TaxedMoneyRange) Equal(other *TaxedMoneyRange) bool {
	return t.Start.Equal(other.Start) && t.Stop.Equal(other.Stop)
}

// LessThan checks if current taxed money range less than given other
func (t *TaxedMoneyRange) LessThan(other *TaxedMoneyRange) bool {
	return t.Start.LessThan(other.Start) && t.Stop.LessThan(other.Stop)
}

// LessThanOrEqual checks if current taxed money range less than or equal to given other
func (t *TaxedMoneyRange) LessThanOrEqual(other *TaxedMoneyRange) bool {
	return t.LessThan(other) || t.Equal(other)
}

// Contains check is given taxed money is in range from start to stop.
//
// start <= item <= stop
func (t *TaxedMoneyRange) Contains(item *TaxedMoney) bool {
	return t.Start.LessThanOrEqual(item) && item.LessThanOrEqual(t.Stop)
}

// Return a copy of the range with start and stop quantized.
// All arguments are passed to `TaxedMoney.quantize` which in turn calls
// `Money.quantize
func (t *TaxedMoneyRange) Quantize(exp *int, round Rounding) (*TaxedMoneyRange, error) {
	start, err := t.Start.Quantize(exp, round)
	if err != nil {
		return nil, err
	}
	stop, err := t.Stop.Quantize(exp, round)
	if err != nil {
		return nil, err
	}
	return &TaxedMoneyRange{
		Start:    start,
		Stop:     stop,
		Currency: t.Currency,
	}, nil
}

// Return a range with start or stop replaced with given values
func (t *TaxedMoneyRange) Replace(start, stop *TaxedMoney) (*TaxedMoneyRange, error) {
	if start == nil {
		start = t.Start
	}
	if stop == nil {
		stop = t.Stop
	}

	return NewTaxedMoneyRange(start, stop)
}

// Apply a fixed discount to TaxedMoneyRange.
func (t *TaxedMoneyRange) FixedDiscount(discount *Money) (*TaxedMoneyRange, error) {
	baseStart, err := t.Start.FixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	baseStop, err := t.Stop.FixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	return NewTaxedMoneyRange(baseStart, baseStop)
}

func (t *TaxedMoneyRange) Mul(other any) (*TaxedMoneyRange, error) {
	panic("not implemented")
}

func (t *TaxedMoneyRange) TrueDiv(other any) (*TaxedMoneyRange, error) {
	panic("not implemented")
}
