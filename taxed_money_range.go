package goprices

import "fmt"

type TaxedMoneyRange struct {
	Start    *TaxedMoney
	Stop     *TaxedMoney
	currency string
}

// NewTaxedMoneyRange create new taxed money range.
// It returns nil and error value if start > stop or they have different currencies
func NewTaxedMoneyRange(start, stop *TaxedMoney) (*TaxedMoneyRange, error) {
	if start.currency != stop.currency {
		return nil, ErrNotSameCurrency
	}
	less, err := stop.LessThan(start)
	if err != nil {
		return nil, err
	}

	if less {
		return nil, ErrStopLessThanStart
	}

	return &TaxedMoneyRange{start, stop, start.currency}, nil
}

// String implements fmt.Stringer interface
func (t *TaxedMoneyRange) String() string {
	return fmt.Sprintf("TaxedMoneyRange{%q, %q}", t.Start.String(), t.Stop.String())
}

// Add adds this taxed money range to a money, money range or taxed money range
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
		return &TaxedMoneyRange{start, stop, t.currency}, nil
	case *MoneyRange:
		start, err := t.Start.Add(v.Start)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Add(v.Stop)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{start, stop, t.currency}, nil
	case *TaxedMoneyRange:
		start, err := t.Start.Add(v.Start)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Add(v.Stop)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{start, stop, t.currency}, nil
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
		return &TaxedMoneyRange{start, stop, t.currency}, nil
	case *MoneyRange:
		start, err := t.Start.Sub(v.Start)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Sub(v.Stop)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{start, stop, t.currency}, nil
	case *TaxedMoneyRange:
		start, err := t.Start.Sub(v.Start)
		if err != nil {
			return nil, err
		}
		stop, err := t.Stop.Sub(v.Stop)
		if err != nil {
			return nil, err
		}
		return &TaxedMoneyRange{start, stop, t.currency}, nil
	default:
		return nil, ErrUnknownType
	}
}

// Equal compares two taxed money range
func (t *TaxedMoneyRange) Equal(other *TaxedMoneyRange) (bool, error) {
	eq1, err := t.Start.Equal(other.Start)
	if err != err {
		return false, err
	}
	eq2, err := t.Stop.Equal(other.Stop)
	if err != err {
		return false, err
	}
	return eq1 && eq2, nil
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

// func (t *TaxedMoneyRange) Quantize() *TaxedMoneyRange {

// }

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
