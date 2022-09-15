package goprices

import (
	"fmt"

	"github.com/site-name/decimal"
)

// MoneyRange has start and stop ends
type MoneyRange struct {
	Start    *Money
	Stop     *Money
	Currency string
}

var (
	_ Currencyable                = (*MoneyRange)(nil)
	_ MoneyInterface[*MoneyRange] = (*MoneyRange)(nil)
)

// NewMoneyRange returns a new range. If start is greater than stop or start and stop have different
// currencies, return nil and non nil error
func NewMoneyRange(start, stop *Money) (*MoneyRange, error) {
	if start == nil || stop == nil {
		return nil, ErrNillValue
	}
	if !start.SameKind(stop) {
		return nil, ErrNotSameCurrency
	}
	_, err := validateCurrency(start.Currency)
	if err != nil {
		return nil, err
	}
	unit, err := validateCurrency(stop.Currency)
	if err != nil {
		return nil, err
	}
	if start.Amount.LessThan(decimal.Zero) || stop.Amount.LessThan(decimal.Zero) {
		return nil, ErrMoneyNegative
	}
	if !start.LessThanOrEqual(stop) {
		return nil, ErrStopLessThanStart
	}

	return &MoneyRange{
		Start:    start,
		Stop:     stop,
		Currency: unit,
	}, nil
}

// String implements fmt.Stringer interface{}
func (m *MoneyRange) String() string {
	return fmt.Sprintf("MoneyRange{%s, %s}", m.Start.String(), m.Stop.String())
}

// MyCurrency returns current money range's Currency
func (m *MoneyRange) MyCurrency() string {
	return m.Currency
}

// Add adds a Value to current.
//
// other must be either *Money or *MoneyRange
func (m *MoneyRange) Add(other interface{}) (*MoneyRange, error) {
	if other == nil {
		return nil, ErrNillValue
	}

	switch v := other.(type) {
	case *Money:
		start, err := m.Start.Add(v)
		if err != nil {
			return nil, err
		}
		stop, err := m.Stop.Add(v)
		if err != nil {
			return nil, err
		}
		return &MoneyRange{start, stop, m.Currency}, nil

	case *MoneyRange:
		start, err := m.Start.Add(v.Start)
		if err != nil {
			return nil, err
		}
		stop, err := m.Stop.Add(v.Stop)
		if err != nil {
			return nil, err
		}
		return &MoneyRange{start, stop, m.Currency}, nil

	default:
		return nil, ErrUnknownType
	}
}

// Sub subtracts current money to given `other`.
// `other` can be either `*Money` or `*MoneyRange`
func (m *MoneyRange) Sub(other interface{}) (*MoneyRange, error) {
	if other == nil {
		return nil, ErrNillValue
	}

	switch v := other.(type) {
	case *Money:
		start, err := m.Start.Sub(v)
		if err != nil {
			return nil, err
		}
		stop, err := m.Stop.Sub(v)
		if err != nil {
			return nil, err
		}
		return &MoneyRange{start, stop, m.Currency}, nil

	case *MoneyRange:
		start, err := m.Start.Sub(v.Start)
		if err != nil {
			return nil, err
		}
		stop, err := m.Stop.Sub(v.Stop)
		if err != nil {
			return nil, err
		}
		return &MoneyRange{start, stop, m.Currency}, nil

	default:
		return nil, ErrUnknownType
	}
}

// Equal Checks if two MoneyRange are equal both `Start`, `Stop` and `Currency`
func (m *MoneyRange) Equal(other *MoneyRange) bool {
	return m.Start.Equal(other.Start) && m.Stop.Equal(other.Stop)
}

// LessThan compares currenct money range to given other
func (m *MoneyRange) LessThan(other *MoneyRange) bool {
	return m.Start.LessThan(other.Start) && m.Stop.LessThan(other.Stop)
}

// LessThanOrEqual checks if current money range is less than or equal given other
func (m *MoneyRange) LessThanOrEqual(other *MoneyRange) bool {
	return m.LessThan(other) || m.Equal(other)
}

// Contains check if a Money is between this MoneyRange's two ends
func (m *MoneyRange) Contains(value *Money) bool {
	return m.Start.LessThanOrEqual(value) && value.LessThanOrEqual(m.Stop)
}

// Return a copy of the range with start and stop quantized.
// All arguments are passed to `Money.quantize
func (m *MoneyRange) Quantize(exp *int, round Rounding) (*MoneyRange, error) {
	start, err := m.Start.Quantize(exp, round)
	if err != nil {
		return nil, err
	}
	stop, err := m.Stop.Quantize(exp, round)
	if err != nil {
		return nil, err
	}
	return &MoneyRange{
		Start:    start,
		Stop:     stop,
		Currency: m.Currency,
	}, nil
}

// Replace replace Start and Stop of currenct MoneyRagne With two given `start` and `stop` respectively.
func (m *MoneyRange) Replace(start, stop *Money) (*MoneyRange, error) {
	if start == nil {
		start = m.Start
	}
	if stop == nil {
		stop = m.Stop
	}
	return NewMoneyRange(start, stop)
}

// Apply a fixed discount to MoneyRange.
func (m *MoneyRange) fixedDiscount(discount *Money) (*MoneyRange, error) {
	baseStart, err := m.Start.fixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	baseStop, err := m.Stop.fixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	return NewMoneyRange(baseStart, baseStop)
}

func (m *MoneyRange) Mul(other any) (*MoneyRange, error) {
	panic("not implemented")
}

func (m *MoneyRange) TrueDiv(other any) (*MoneyRange, error) {
	panic("not implemented")
}

func (m *MoneyRange) fractionalDiscount(fraction decimal.Decimal, fromGross bool) (*MoneyRange, error) {
	start, err1 := m.Start.fractionalDiscount(fraction, fromGross)
	if err1 != nil {
		return nil, err1
	}

	stop, err2 := m.Stop.fractionalDiscount(fraction, fromGross)
	if err2 != nil {
		return nil, err2
	}

	return NewMoneyRange(start, stop)
}
