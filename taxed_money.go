package goprices

import "fmt"

type TaxedMoney struct {
	Net      *Money
	Gross    *Money
	currency string
}

func NewTaxedMoney(net, gross *Money) (*TaxedMoney, error) {
	if err := net.sameKind(gross); err != nil {
		return nil, err
	}

	return &TaxedMoney{net, gross, net.Currency}, nil
}

func (t *TaxedMoney) String() string {
	return fmt.Sprintf("TaxedMoney{net=%q, gross=%q}", t.Net.String(), t.Gross.String())
}

func (t *TaxedMoney) LessThan(other *TaxedMoney) (bool, error) {
	if other.currency != t.currency {
		return false, ErrNotSameCurrency
	}
	return t.Gross.LessThan(other.Gross)
}

func (t *TaxedMoney) Equal(other *TaxedMoney) (bool, error) {
	if t.currency != other.currency {
		return false, ErrNotSameCurrency
	}

	eq1, _ := t.Net.Equal(other.Net)
	eq2, _ := t.Gross.Equal(other.Gross)

	return eq1 && eq2, nil
}

func (t *TaxedMoney) TrueDiv(other *TaxedMoney) (*TaxedMoney, error) {
	if t.currency != other.currency {
		return nil, ErrNotSameCurrency
	}
	net, _ := t.Net.TrueDiv(other.Net)
	gross, _ := t.Gross.TrueDiv(other.Gross)
	return NewTaxedMoney(net, gross)
}

func (t *TaxedMoney) Add(other interface{}) (*TaxedMoney, error) {
	switch v := other.(type) {
	case *Money:
		if v.Currency != t.currency {
			return nil, ErrNotSameCurrency
		}
		net, _ := t.Net.Add(v)
		gross, _ := t.Gross.Add(v)
		return NewTaxedMoney(net, gross)
	case *TaxedMoney:
		if t.currency != v.currency {
			return nil, ErrNotSameCurrency
		}
		net, _ := t.Net.Add(v.Net)
		gross, _ := t.Gross.Add(v.Gross)
		return NewTaxedMoney(net, gross)
	default:
		return nil, ErrUnknownType
	}
}

func (t *TaxedMoney) Sub(other interface{}) (*TaxedMoney, error) {
	switch v := other.(type) {
	case *Money:
		if v.Currency != t.currency {
			return nil, ErrNotSameCurrency
		}
		net, _ := t.Net.Sub(v)
		gross, _ := t.Gross.Sub(v)
		return NewTaxedMoney(net, gross)
	case *TaxedMoney:
		if t.currency != v.currency {
			return nil, ErrNotSameCurrency
		}
		net, _ := t.Net.Sub(v.Net)
		gross, _ := t.Gross.Sub(v.Gross)
		return NewTaxedMoney(net, gross)
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
