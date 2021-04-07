package goprices

import "fmt"

// TaxedMoney represents taxed money. It wraps net and gross
type TaxedMoney struct {
	Net      *Money
	Gross    *Money
	currency string
}

// NewTaxedMoney returns new TaxedMoney,
// If net and gross have different currency type, return nil and error
func NewTaxedMoney(net, gross *Money) (*TaxedMoney, error) {
	if err := net.sameKind(gross); err != nil {
		return nil, err
	}

	return &TaxedMoney{net, gross, net.Currency}, nil
}

// String implements fmt.Stringer interface
func (t *TaxedMoney) String() string {
	return fmt.Sprintf("TaxedMoney{net=%q, gross=%q}", t.Net.String(), t.Gross.String())
}

// LessThan compare this to other TaxedMoney, return false and error if operants have different currencies
func (t *TaxedMoney) LessThan(other *TaxedMoney) (bool, error) {
	return t.Gross.LessThan(other.Gross) // currency type check included
}

// Equal checks if two taxed money euqal both net and gross
func (t *TaxedMoney) Equal(other *TaxedMoney) (bool, error) {
	eq1, err := t.Net.Equal(other.Net)
	if err != nil {
		return false, err
	}
	eq2, err := t.Gross.Equal(other.Gross)
	if err != nil {
		return false, err
	}

	return eq1 && eq2, nil
}

// LessThanOrEqual checks if two taxed money are less than or euqal each other
func (t *TaxedMoney) LessThanOrEqual(other *TaxedMoney) (bool, error) {
	less, err := t.LessThan(other)
	if err != nil {
		return false, err
	}
	eq, err := t.Equal(other)
	if err != nil {
		return false, err
	}
	return less || eq, nil
}

// TrueDiv divides two taxed money
func (t *TaxedMoney) TrueDiv(other *TaxedMoney) (*TaxedMoney, error) {
	net, err := t.Net.TrueDiv(other.Net)
	if err != nil {
		return nil, err
	}
	gross, err := t.Gross.TrueDiv(other.Gross)
	if err != nil {
		return nil, err
	}
	return &TaxedMoney{net, gross, t.currency}, nil
}

// Add adds two taxed money, or taxed money to a money.
func (t *TaxedMoney) Add(other interface{}) (*TaxedMoney, error) {
	switch v := other.(type) {
	case *Money:
		net, err := t.Net.Add(v)
		if err != nil {
			return nil, err
		}
		gross, err := t.Gross.Add(v)
		if err != nil {
			return nil, err
		}
		return &TaxedMoney{net, gross, t.currency}, nil
	case *TaxedMoney:
		net, err := t.Net.Add(v.Net)
		if err != nil {
			return nil, err
		}
		gross, err := t.Gross.Add(v.Gross)
		if err != nil {
			return nil, err
		}
		return &TaxedMoney{net, gross, t.currency}, nil
	default:
		return nil, ErrUnknownType
	}
}

// Add substract two taxed money, or taxed money to a money.
func (t *TaxedMoney) Sub(other interface{}) (*TaxedMoney, error) {
	switch v := other.(type) {
	case *Money:
		net, err := t.Net.Sub(v)
		if err != nil {
			return nil, err
		}
		gross, err := t.Gross.Sub(v)
		if err != nil {
			return nil, err
		}
		return &TaxedMoney{net, gross, t.currency}, nil
	case *TaxedMoney:
		net, err := t.Net.Sub(v.Net)
		if err != nil {
			return nil, err
		}
		gross, err := t.Gross.Sub(v.Gross)
		if err != nil {
			return nil, err
		}
		return &TaxedMoney{net, gross, t.currency}, nil
	default:
		return nil, ErrUnknownType
	}
}

func (t *TaxedMoney) Tax() (*Money, error) {
	return t.Gross.Sub(t.Net)
}

// func (t *TaxedMoney) Quantize() (*TaxedMoney, error) {
//
// }
