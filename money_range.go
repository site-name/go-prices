package goprices

import (
	"fmt"

	"github.com/site-name/decimal"
)

// MoneyRange has start and stop ends
type MoneyRange struct {
	start Money
	stop  Money
}

var _ MoneyInterface[MoneyRange] = (*MoneyRange)(nil)

// NewMoneyRange returns a new range. If start is greater than stop or start and stop have different
// currencies, return nil and non nil error
func NewMoneyRange(start, stop Money) (*MoneyRange, error) {
	startUnit, err := validateCurrency(start.currency)
	if err != nil {
		return nil, err
	}
	stopUnit, err := validateCurrency(stop.currency)
	if err != nil {
		return nil, err
	}
	if startUnit != stopUnit {
		return nil, ErrNotSameCurrency
	}
	if start.amount.LessThan(decimal.Zero) || stop.amount.LessThan(decimal.Zero) {
		return nil, ErrMoneyNegative
	}
	if stop.LessThanOrEqual(start) {
		return nil, ErrStopLessThanStart
	}

	return &MoneyRange{
		start: start,
		stop:  stop,
	}, nil
}

func NewMoneyRangeFromFloats(start, stop float64, currency string) (*MoneyRange, error) {
	startMoney, err := NewMoney(start, currency)
	if err != nil {
		return nil, err
	}
	stopMoney, err := NewMoney(stop, currency)
	if err != nil {
		return nil, err
	}
	return NewMoneyRange(*startMoney, *stopMoney)
}

func NewMoneyRangeFromDecimals(start, stop decimal.Decimal, currency string) (*MoneyRange, error) {
	startMoney, err := NewMoneyFromDecimal(start, currency)
	if err != nil {
		return nil, err
	}
	stopMoney, err := NewMoneyFromDecimal(stop, currency)
	if err != nil {
		return nil, err
	}
	return NewMoneyRange(*startMoney, *stopMoney)
}

func (m MoneyRange) GetStart() Money {
	return m.start
}

func (m *MoneyRange) SetStart(start Money) {
	if m == nil {
		return
	}
	m.start = start
}

func (m *MoneyRange) SetStop(stop Money) {
	if m == nil {
		return
	}
	m.stop = stop
}

func (m MoneyRange) GetStop() Money {
	return m.stop
}

func (m MoneyRange) String() string {
	return fmt.Sprintf("MoneyRange{%s, %s}", m.start.String(), m.stop.String())
}

// GetCurrency returns current money range's currency
func (m MoneyRange) GetCurrency() string {
	return m.start.currency
}

// Add adds a Value to current.
//
// other must be either Money or MoneyRange
func (m MoneyRange) Add(other any) (*MoneyRange, error) {
	if other == nil {
		return nil, ErrNillValue
	}

	switch v := other.(type) {
	case Money:
		start, err := m.start.Add(v)
		if err != nil {
			return nil, err
		}
		stop, err := m.stop.Add(v)
		if err != nil {
			return nil, err
		}
		return &MoneyRange{*start, *stop}, nil

	case MoneyRange:
		start, err := m.start.Add(v.start)
		if err != nil {
			return nil, err
		}
		stop, err := m.stop.Add(v.stop)
		if err != nil {
			return nil, err
		}
		return &MoneyRange{*start, *stop}, nil

	default:
		return nil, ErrUnknownType
	}
}

// Sub subtracts current money to given `other`.
// `other` can be either `Money` or `MoneyRange`
func (m MoneyRange) Sub(other any) (*MoneyRange, error) {
	if other == nil {
		return nil, ErrNillValue
	}

	switch v := other.(type) {
	case Money:
		return m.Add(v.Neg())
	case MoneyRange:
		return m.Add(v.Neg())

	default:
		return nil, ErrUnknownType
	}
}

func (m MoneyRange) Neg() MoneyRange {
	return MoneyRange{
		start: m.start.Neg(),
		stop:  m.stop.Neg(),
	}
}

// Equal Checks if two MoneyRange are equal both `start`, `stop` and `currency`
func (m MoneyRange) Equal(other MoneyRange) bool {
	return m.start.Equal(other.start) && m.stop.Equal(other.stop)
}

// LessThan compares currenct money range to given other
func (m MoneyRange) LessThan(other MoneyRange) bool {
	return m.start.LessThan(other.start) && m.stop.LessThan(other.stop)
}

// LessThanOrEqual checks if current money range is less than or equal given other
func (m MoneyRange) LessThanOrEqual(other MoneyRange) bool {
	return m.LessThan(other) || m.Equal(other)
}

// Contains check if a Money is between this MoneyRange's two ends
func (m MoneyRange) Contains(value Money) bool {
	return m.start.LessThanOrEqual(value) && value.LessThanOrEqual(m.stop)
}

// Return a copy of the range with start and stop quantized.
// NOTE: if exp < 0 the system will use default
func (m MoneyRange) Quantize(round Rounding, exp int) (*MoneyRange, error) {
	start, err := m.start.Quantize(round, exp)
	if err != nil {
		return nil, err
	}
	stop, err := m.stop.Quantize(round, exp)
	if err != nil {
		return nil, err
	}
	return &MoneyRange{
		start: *start,
		stop:  *stop,
	}, nil
}

// Replace replace start and stop of currenct MoneyRagne With two given `start` and `stop` respectively.
func (m MoneyRange) Replace(start, stop *Money) (*MoneyRange, error) {
	if start == nil {
		start = &m.start
	}
	if stop == nil {
		stop = &m.stop
	}
	return NewMoneyRange(*start, *stop)
}

// Apply a fixed discount to MoneyRange.
func (m MoneyRange) fixedDiscount(discount Money) (*MoneyRange, error) {
	baseStart, err := m.start.fixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	baseStop, err := m.stop.fixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	return NewMoneyRange(*baseStart, *baseStop)
}

func (m MoneyRange) Mul(other float64) MoneyRange {
	return MoneyRange{
		start: m.start.Mul(other),
		stop:  m.start.Mul(other),
	}
}

func (m MoneyRange) TrueDiv(other float64) MoneyRange {
	return MoneyRange{
		start: m.start.TrueDiv(other),
		stop:  m.start.TrueDiv(other),
	}
}

func (m MoneyRange) fractionalDiscount(fraction decimal.Decimal, fromGross bool, rounding Rounding) (*MoneyRange, error) {
	start, err1 := m.start.fractionalDiscount(fraction, fromGross, rounding)
	if err1 != nil {
		return nil, err1
	}

	stop, err2 := m.stop.fractionalDiscount(fraction, fromGross, rounding)
	if err2 != nil {
		return nil, err2
	}

	return NewMoneyRange(*start, *stop)
}
