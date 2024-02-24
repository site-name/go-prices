package goprices

import (
	"fmt"

	"github.com/site-name/decimal"
)

type TaxedMoneyRange struct {
	Start TaxedMoney
	Stop  TaxedMoney
}

var (
	_ Currencier                       = (*TaxedMoneyRange)(nil)
	_ MoneyInterface[*TaxedMoneyRange] = (*TaxedMoneyRange)(nil)
)

// NewTaxedMoneyRange create new taxed money range.
// It returns nil and error value if start > stop or they have different currencies
func NewTaxedMoneyRange(start, stop TaxedMoney) (*TaxedMoneyRange, error) {
	startUnit, err := validateCurrency(start.GetCurrency())
	if err != nil {
		return nil, err
	}
	stopUnit, err := validateCurrency(stop.GetCurrency())
	if err != nil {
		return nil, err
	}
	if startUnit != stopUnit {
		return nil, ErrNotSameCurrency
	}
	if start.Net.Amount.LessThan(decimal.Zero) || stop.Gross.Amount.LessThan(decimal.Zero) {
		return nil, ErrMoneyNegative
	}

	if stop.LessThan(start) {
		return nil, ErrStopLessThanStart
	}

	return &TaxedMoneyRange{start, stop}, nil
}

// String implements fmt.Stringer interface
func (t *TaxedMoneyRange) String() string {
	return fmt.Sprintf("TaxedMoneyRange{%s, %s}", t.Start.String(), t.Stop.String())
}

// GetCurrency returns current taxed money range's Currency
func (t *TaxedMoneyRange) GetCurrency() string {
	return t.Start.GetCurrency()
}

// Add adds this taxed money range to another value
// other must be either: Money, MoneyRange or TaxedMoneyRange or TaxedMoney
func (t *TaxedMoneyRange) Add(other any) (*TaxedMoneyRange, error) {
	switch v := other.(type) {
	case Money, TaxedMoney:
		start, err := t.Start.Add(v)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Add(v)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{*start, *stop}, nil

	case MoneyRange:
		start, err := t.Start.Add(v.Start)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Add(v.Stop)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{*start, *stop}, nil

	case TaxedMoneyRange:
		start, err := t.Start.Add(v.Start)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Add(v.Stop)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{*start, *stop}, nil

	default:
		return nil, ErrUnknownType
	}
}

// Sub substract this taxed money range to given other.
// other must be either Money or TaxedMoney or MoneyRange or TaxedMoneyRange
func (t *TaxedMoneyRange) Sub(other any) (*TaxedMoneyRange, error) {
	if other == nil {
		return nil, ErrNillValue
	}

	switch v := other.(type) {
	case Money, TaxedMoney:
		start, err := t.Start.Sub(v)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Sub(v)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{*start, *stop}, nil

	case MoneyRange:
		start, err := t.Start.Sub(v.Start)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Sub(v.Stop)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{*start, *stop}, nil

	case TaxedMoneyRange:
		start, err := t.Start.Sub(v.Start)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Sub(v.Stop)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{*start, *stop}, nil

	default:
		return nil, ErrUnknownType
	}
}

// Equal compares two taxed money range
func (t *TaxedMoneyRange) Equal(other TaxedMoneyRange) bool {
	return t.Start.Equal(other.Start) && t.Stop.Equal(other.Stop)
}

// LessThan checks if current taxed money range less than given other
func (t *TaxedMoneyRange) LessThan(other TaxedMoneyRange) bool {
	return t.Start.LessThan(other.Start) && t.Stop.LessThan(other.Stop)
}

// LessThanOrEqual checks if current taxed money range less than or equal to given other
func (t *TaxedMoneyRange) LessThanOrEqual(other TaxedMoneyRange) bool {
	return t.LessThan(other) || t.Equal(other)
}

// Contains check is given taxed money is in range from start to stop.
//
// start <= item <= stop
func (t *TaxedMoneyRange) Contains(item TaxedMoney) bool {
	return t.Start.LessThanOrEqual(item) && item.LessThanOrEqual(t.Stop)
}

// Return a copy of the range with start and stop quantized.
// NOTE: if exp < 0; default will be used
func (t *TaxedMoneyRange) Quantize(round Rounding, exp int) (*TaxedMoneyRange, error) {
	start, err := t.Start.Quantize(round, exp)
	if err != nil {
		return nil, err
	}
	stop, err := t.Stop.Quantize(round, exp)
	if err != nil {
		return nil, err
	}
	return &TaxedMoneyRange{
		Start: *start,
		Stop:  *stop,
	}, nil
}

// Return a range with start or stop replaced with given values
func (t *TaxedMoneyRange) Replace(start, stop *TaxedMoney) (*TaxedMoneyRange, error) {
	if start == nil {
		start = &t.Start
	}
	if stop == nil {
		stop = &t.Stop
	}

	return NewTaxedMoneyRange(*start, *stop)
}

// Apply a fixed discount to TaxedMoneyRange.
func (t *TaxedMoneyRange) fixedDiscount(discount Money) (*TaxedMoneyRange, error) {
	baseStart, err := t.Start.fixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	baseStop, err := t.Stop.fixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	return NewTaxedMoneyRange(*baseStart, *baseStop)
}

func (t *TaxedMoneyRange) Mul(other float64) TaxedMoneyRange {
	return TaxedMoneyRange{
		Start: t.Start.Mul(other),
		Stop:  t.Stop.Mul(other),
	}
}

func (t *TaxedMoneyRange) TrueDiv(other float64) *TaxedMoneyRange {
	return &TaxedMoneyRange{
		Start: t.Start.TrueDiv(other),
		Stop:  t.Stop.TrueDiv(other),
	}
}

func (m *TaxedMoneyRange) fractionalDiscount(fraction decimal.Decimal, fromGross bool) (*TaxedMoneyRange, error) {
	start, err := m.Start.fractionalDiscount(fraction, fromGross)
	if err != nil {
		return nil, err
	}

	stop, err := m.Stop.fractionalDiscount(fraction, fromGross)
	if err != nil {
		return nil, err
	}

	return NewTaxedMoneyRange(*start, *stop)
}
