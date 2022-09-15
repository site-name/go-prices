package goprices

import (
	"fmt"

	"github.com/site-name/decimal"
)

// TaxedMoney represents taxed money. It wraps net, gross money and currency.
type TaxedMoney struct {
	Net      *Money
	Gross    *Money
	Currency string
}

var (
	_ Currencyable                = (*TaxedMoney)(nil)
	_ MoneyInterface[*TaxedMoney] = (*TaxedMoney)(nil)
)

// NewTaxedMoney returns new TaxedMoney,
// If net and gross have different currency type, return nil and error
func NewTaxedMoney(net, gross *Money) (*TaxedMoney, error) {
	if net == nil || gross == nil {
		return nil, ErrNillValue
	}
	unit1, err := validateCurrency(net.Currency)
	if err != nil {
		return nil, err
	}
	unit2, err := validateCurrency(gross.Currency)
	if err != nil {
		return nil, err
	}

	if unit1 != unit2 {
		return nil, ErrNotSameCurrency
	}

	return &TaxedMoney{net, gross, unit1}, nil
}

// String implements fmt.Stringer interface
func (t *TaxedMoney) String() string {
	return fmt.Sprintf("TaxedMoney{net=%s, gross=%s}", t.Net.String(), t.Gross.String())
}

// MyCurrency returns current taxed money's Currency
func (m *TaxedMoney) MyCurrency() string {
	return m.Currency
}

// LessThan check if this money's gross is less than other's gross
func (t *TaxedMoney) LessThan(other *TaxedMoney) bool {
	return t.Gross.LessThan(other.Gross)
}

// Equal checks if two taxed money are equal both in net and gross
func (t *TaxedMoney) Equal(other *TaxedMoney) bool {
	return t.Net.Equal(other.Net) && t.Gross.Equal(other.Gross)
}

// LessThanOrEqual checks if this money is less than or equal to other.
func (t *TaxedMoney) LessThanOrEqual(other *TaxedMoney) bool {
	return t.LessThan(other) || t.Equal(other)
}

// Mul multiplies current taxed money with given other
//
// other must only be either ints or floats or Decimal
func (m *TaxedMoney) Mul(other interface{}) (*TaxedMoney, error) {
	if other == nil {
		return nil, ErrNillValue
	}

	net, err := m.Net.Mul(other)
	if err != nil {
		return nil, err
	}
	gross, err := m.Gross.Mul(other)
	if err != nil {
		return nil, err
	}

	return &TaxedMoney{
		Net:      net,
		Gross:    gross,
		Currency: m.Currency,
	}, nil
}

// TrueDiv divides current tabled money to other.
// other must be either Decimal or ints or floats
func (t *TaxedMoney) TrueDiv(other interface{}) (*TaxedMoney, error) {
	var (
		newNet   *Money
		newGross *Money
		err      error
	)

	newNet, err = t.Net.TrueDiv(other)
	if err != nil {
		return nil, err
	}
	newGross, err = t.Gross.TrueDiv(other)
	if err != nil {
		return nil, err
	}

	return &TaxedMoney{
		Gross:    newGross,
		Net:      newNet,
		Currency: t.Currency,
	}, nil
}

// Add adds a money or taxed money to this.
// other must be either *Money or *TaxedMoney
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
		return &TaxedMoney{net, gross, t.Currency}, nil

	case *TaxedMoney:
		net, err := t.Net.Add(v.Net)
		if err != nil {
			return nil, err
		}
		gross, err := t.Gross.Add(v.Gross)
		if err != nil {
			return nil, err
		}
		return &TaxedMoney{net, gross, t.Currency}, nil

	default:
		return nil, ErrUnknownType
	}
}

// Add substract this money to other.
// other must be either *Money or *TaxedMoney.
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
		return &TaxedMoney{net, gross, t.Currency}, nil

	case *TaxedMoney:
		net, err := t.Net.Sub(v.Net)
		if err != nil {
			return nil, err
		}
		gross, err := t.Gross.Sub(v.Gross)
		if err != nil {
			return nil, err
		}
		return &TaxedMoney{net, gross, t.Currency}, nil

	default:
		return nil, ErrUnknownType
	}
}

// Tax calculates taxed money by subtracting m's gross to m's net
func (t *TaxedMoney) Tax() (*Money, error) {
	return t.Gross.Sub(t.Net)
}

// Return a new instance with both net and gross quantized.
// All arguments are passed to `Money.quantize
func (t *TaxedMoney) Quantize(exp *int, round Rounding) (*TaxedMoney, error) {
	net, err := t.Net.Quantize(exp, round)
	if err != nil {
		return nil, err
	}
	gross, err := t.Gross.Quantize(exp, round)
	if err != nil {
		return nil, err
	}

	return &TaxedMoney{
		Net:      net,
		Gross:    gross,
		Currency: t.Currency,
	}, nil
}

// Apply a fixed discount to TaxedMoney.
func (t *TaxedMoney) fixedDiscount(discount *Money) (*TaxedMoney, error) {
	baseNet, err := t.Net.fixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	baseGross, err := t.Gross.fixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	return NewTaxedMoney(baseNet, baseGross)
}

func (m *TaxedMoney) fractionalDiscount(fraction decimal.Decimal, fromGross bool) (*TaxedMoney, error) {
	op := &Money{
		Currency: m.Currency,
		Amount:   m.Gross.Amount,
	}
	if !fromGross {
		op.Amount = m.Net.Amount
	}

	op, _ = op.Mul(fraction)
	discount, err := op.Quantize(nil, Down)
	if err != nil {
		return nil, err
	}

	return m.fixedDiscount(discount)
}
