package goprices

import (
	"fmt"
)

type TaxedMoneyRange struct {
	Start    *TaxedMoney
	Stop     *TaxedMoney
	Currency string
}

// NewTaxedMoneyRange create new taxed money range.
// It returns nil and error value if start > stop or they have different currencies
func NewTaxedMoneyRange(start, stop *TaxedMoney) (*TaxedMoneyRange, error) {
	_, err := checkCurrency(start.Currency)
	if err != nil {
		return nil, err
	}
	unit, err := checkCurrency(stop.Currency)
	if err != nil {
		return nil, err
	}

	less, err := stop.LessThan(start)
	if err != nil {
		return nil, err
	}
	if less {
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

// Sub substract this taxed money range to a money, money range or taxed money range
func (t *TaxedMoneyRange) Sub(other interface{}) (*TaxedMoneyRange, error) {
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
func (t *TaxedMoneyRange) Equal(other *TaxedMoneyRange) (bool, error) {
	eq1, err := t.Start.Equal(other.Start)
	if err != nil {
		return false, err
	}
	eq2, err := t.Stop.Equal(other.Stop)
	if err != nil {
		return false, err
	}
	return eq1 && eq2, nil
}

// LessThan checks if current taxed money range less than given other
func (t *TaxedMoneyRange) LessThan(other *TaxedMoneyRange) (bool, error) {
	l1, err := t.Start.LessThan(other.Start)
	if err != nil {
		return false, err
	}
	l2, err := t.Stop.LessThan(other.Stop)
	if err != nil {
		return false, err
	}

	return l1 && l2, nil
}

// LessThanOrEqual checks if current taxed money range less than or equal to given other
func (t *TaxedMoneyRange) LessThanOrEqual(other *TaxedMoneyRange) (bool, error) {
	less, err := t.LessThan(other)
	if err != nil {
		return false, err
	}
	equal, err := t.Equal(other)
	if err != nil {
		return false, err
	}

	return less || equal, nil
}

// Contains check is given taxed money is in range from start to stop.
//
//start <= item <= stop
func (t *TaxedMoneyRange) Contains(item *TaxedMoney) (bool, error) {
	greaterThanStart, err := t.Start.LessThanOrEqual(item)
	if err != nil {
		return false, err
	}
	lessThanStop, err := item.LessThanOrEqual(t.Stop)
	if err != nil {
		return false, err
	}
	return greaterThanStart && lessThanStop, nil
}

// Return a copy of the range with start and stop quantized.
// All arguments are passed to `TaxedMoney.quantize` which in turn calls
// `Money.quantize
func (t *TaxedMoneyRange) Quantize(exp *int32, round Rounding) (*TaxedMoneyRange, error) {
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
