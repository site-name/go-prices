package goprices

import (
	"fmt"

	"github.com/site-name/decimal"
)

type TaxedMoneyRange struct {
	start TaxedMoney
	stop  TaxedMoney
}

var _ MoneyInterface[TaxedMoneyRange] = (*TaxedMoneyRange)(nil)

func (t *TaxedMoneyRange) GetStart() TaxedMoney {
	if t == nil {
		panic(ErrNillValue)
	}
	return t.start
}

func (t *TaxedMoneyRange) GetStop() TaxedMoney {
	if t == nil {
		panic(ErrNillValue)
	}
	return t.stop
}

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
	if start.net.amount.LessThan(decimal.Zero) || stop.gross.amount.LessThan(decimal.Zero) {
		return nil, ErrMoneyNegative
	}

	if stop.LessThan(start) {
		return nil, ErrStopLessThanStart
	}

	return &TaxedMoneyRange{start, stop}, nil
}

func (t *TaxedMoneyRange) Neg() TaxedMoneyRange {
	if t == nil {
		panic(ErrNillValue)
	}
	return TaxedMoneyRange{
		start: t.start,
		stop:  t.stop,
	}
}

// String implements fmt.Stringer interface
func (t *TaxedMoneyRange) String() string {
	if t == nil {
		panic(ErrNillValue)
	}
	return fmt.Sprintf("TaxedMoneyRange{%s, %s}", t.start.String(), t.stop.String())
}

func (m *TaxedMoneyRange) SetStart(start TaxedMoney) {
	m.start = start
}

func (m *TaxedMoneyRange) SetStop(stop TaxedMoney) {
	m.stop = stop
}

// GetCurrency returns current taxed money range's Currency
func (t *TaxedMoneyRange) GetCurrency() string {
	if t == nil {
		panic(ErrNillValue)
	}
	return t.start.gross.currency
}

// Add adds this taxed money range to another value
// other must be either: Money, MoneyRange or TaxedMoneyRange or TaxedMoney
func (t *TaxedMoneyRange) Add(other any) (*TaxedMoneyRange, error) {
	if t == nil || other == nil {
		return nil, ErrNillValue
	}

	switch v := other.(type) {
	case Money, TaxedMoney:
		start, err := t.start.Add(v)
		if err != nil {
			return nil, err
		}
		stop, err := t.stop.Add(v)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{*start, *stop}, nil

	case MoneyRange:
		start, err := t.start.Add(v.start)
		if err != nil {
			return nil, err
		}
		stop, err := t.stop.Add(v.stop)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{*start, *stop}, nil

	case TaxedMoneyRange:
		start, err := t.start.Add(v.start)
		if err != nil {
			return nil, err
		}
		stop, err := t.stop.Add(v.stop)
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
	if other == nil || t == nil {
		return nil, ErrNillValue
	}

	switch v := other.(type) {
	case Money:
		return t.Add(v.Neg())
	case TaxedMoney:
		return t.Add(v.Neg())
	case MoneyRange:
		return t.Add(v.Neg())
	case TaxedMoneyRange:
		return t.Add(v.Neg())

	default:
		return nil, ErrUnknownType
	}
}

// Equal compares two taxed money range
func (t *TaxedMoneyRange) Equal(other TaxedMoneyRange) bool {
	return t != nil && t.start.Equal(other.start) && t.stop.Equal(other.stop)
}

// LessThan checks if current taxed money range less than given other
func (t *TaxedMoneyRange) LessThan(other TaxedMoneyRange) bool {
	return t != nil && t.start.LessThan(other.start) && t.stop.LessThan(other.stop)
}

// LessThanOrEqual checks if current taxed money range less than or equal to given other
func (t *TaxedMoneyRange) LessThanOrEqual(other TaxedMoneyRange) bool {
	return t != nil && t.LessThan(other) || t.Equal(other)
}

// Contains check is given taxed money is in range from start to stop.
//
// start <= item <= stop
func (t *TaxedMoneyRange) Contains(item TaxedMoney) bool {
	return t != nil && t.start.LessThanOrEqual(item) && item.LessThanOrEqual(t.stop)
}

// Return a copy of the range with start and stop quantized.
// NOTE: if exp < 0; default will be used
func (t *TaxedMoneyRange) Quantize(round Rounding, exp int) (*TaxedMoneyRange, error) {
	if t == nil {
		return nil, ErrNillValue
	}

	start, err := t.start.Quantize(round, exp)
	if err != nil {
		return nil, err
	}
	stop, err := t.stop.Quantize(round, exp)
	if err != nil {
		return nil, err
	}
	return &TaxedMoneyRange{
		start: *start,
		stop:  *stop,
	}, nil
}

// Return a range with start or stop replaced with given values
func (t *TaxedMoneyRange) Replace(start, stop *TaxedMoney) (*TaxedMoneyRange, error) {
	if t == nil {
		return nil, ErrNillValue
	}
	if start == nil {
		start = &t.start
	}
	if stop == nil {
		stop = &t.stop
	}

	return NewTaxedMoneyRange(*start, *stop)
}

// Apply a fixed discount to TaxedMoneyRange.
func (t *TaxedMoneyRange) fixedDiscount(discount Money) (*TaxedMoneyRange, error) {
	if t == nil {
		return nil, ErrNillValue
	}

	baseStart, err := t.start.fixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	baseStop, err := t.stop.fixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	return NewTaxedMoneyRange(*baseStart, *baseStop)
}

func (t *TaxedMoneyRange) Mul(other float64) TaxedMoneyRange {
	if t == nil {
		panic(ErrNillValue)
	}
	return TaxedMoneyRange{
		start: t.start.Mul(other),
		stop:  t.stop.Mul(other),
	}
}

func (t *TaxedMoneyRange) TrueDiv(other float64) TaxedMoneyRange {
	if t == nil {
		panic(ErrNillValue)
	}
	return TaxedMoneyRange{
		start: t.start.TrueDiv(other),
		stop:  t.stop.TrueDiv(other),
	}
}

func (m *TaxedMoneyRange) fractionalDiscount(fraction decimal.Decimal, fromGross bool) (*TaxedMoneyRange, error) {
	if m == nil {
		return nil, ErrNillValue
	}

	start, err := m.start.fractionalDiscount(fraction, fromGross)
	if err != nil {
		return nil, err
	}

	stop, err := m.stop.fractionalDiscount(fraction, fromGross)
	if err != nil {
		return nil, err
	}

	return NewTaxedMoneyRange(*start, *stop)
}
