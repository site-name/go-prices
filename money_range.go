package goprices

import (
	"errors"
	"fmt"
)

var (
	ErrStopLessThanStart = errors.New("stop must be greater than start")
)

// A taxed money range
type MoneyRange struct {
	Start    *Money
	Stop     *Money
	currency string
}

func NewMoneyRange(start, stop *Money) (*MoneyRange, error) {
	if err := start.sameKind(stop); err != nil {
		return nil, err
	}
	if ok, _ := stop.LessThan(start); ok {
		return nil, ErrStopLessThanStart
	}

	return &MoneyRange{
		Start:    start,
		Stop:     stop,
		currency: start.Currency,
	}, nil
}

// String implements fmt.Stringer interface{}
func (m *MoneyRange) String() string {
	return fmt.Sprintf("Money{%q, %q}", m.Start.String(), m.Stop.String())
}

// Add adds a Money or MoneyRange to this MoneyRange
func (m *MoneyRange) Add(other interface{}) (*MoneyRange, error) {
	switch v := other.(type) {
	case *Money:
		if m.currency != v.Currency {
			return nil, ErrNotSameCurrency
		}
		start, _ := m.Start.Add(v) // already checked err above
		stop, _ := m.Start.Add(v)  // already checked err above
		return &MoneyRange{start, stop, m.currency}, nil
	case *MoneyRange:
		if v.Start.Currency != m.currency {
			return nil, ErrNotSameCurrency
		}
		start, _ := m.Start.Add(v.Start) // already checked err above
		stop, _ := m.Stop.Add(v.Stop)    // already checked err above
		return &MoneyRange{start, stop, m.currency}, nil
	default:
		return nil, ErrUnknownType
	}
}

// Sub reduces a Money or MoneyRange to this MoneyRange
func (m *MoneyRange) Sub(other interface{}) (*MoneyRange, error) {
	switch v := other.(type) {
	case *Money:
		if m.currency != v.Currency {
			return nil, ErrNotSameCurrency
		}
		start, _ := m.Start.Sub(v) // already checked err above
		stop, _ := m.Start.Sub(v)  // already checked err above
		return &MoneyRange{start, stop, m.currency}, nil
	case *MoneyRange:
		if v.Start.Currency != m.currency {
			return nil, ErrNotSameCurrency
		}
		start, _ := m.Start.Sub(v.Start) // already checked err above
		stop, _ := m.Stop.Sub(v.Stop)    // already checked err above
		return &MoneyRange{start, stop, m.currency}, nil
	default:
		return nil, ErrUnknownType
	}
}

// Equal Checks if two MoneyRange are equal
func (m *MoneyRange) Equal(other *MoneyRange) bool {
	if m.currency != other.currency {
		return false
	}
	b1, _ := m.Start.Equal(other.Start) // already checked err above
	b2, _ := m.Stop.Equal(other.Stop)   // already checked err above
	return b1 && b2
}

// Contains check if a Money is between this MoneyRange's two ends
func (m *MoneyRange) Contains(item *Money) bool {
	if m.currency != item.Currency {
		return false
	}
	itemGreaterThanStart, _ := m.Start.LessThan(item) // already checked err above
	itemLessThanStop, _ := item.LessThan(m.Stop)      // already checked err above
	return itemGreaterThanStart && itemLessThanStop
}

// func (m *MoneyRange) Quantise() {

// }

func (m *MoneyRange) Replace(start, stop *Money) (*MoneyRange, error) {
	return NewMoneyRange(start, stop)
}
