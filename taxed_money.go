package goprices

import (
	"fmt"

	"github.com/site-name/decimal"
)

// TaxedMoney represents taxed money. It wraps net, gross money and currency.
type TaxedMoney struct {
	Net   Money
	Gross Money
}

var (
	_ Currencier                  = (*TaxedMoney)(nil)
	_ MoneyInterface[*TaxedMoney] = (*TaxedMoney)(nil)
)

// NewTaxedMoney returns new TaxedMoney,
// If net and gross have different currency type, return nil and error
func NewTaxedMoney(net, gross Money) (*TaxedMoney, error) {
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

	return &TaxedMoney{net, gross}, nil
}

// String implements fmt.Stringer interface
func (t *TaxedMoney) String() string {
	return fmt.Sprintf("TaxedMoney{net=%s, gross=%s}", t.Net.String(), t.Gross.String())
}

// GetCurrency returns current taxed money's Currency
func (m *TaxedMoney) GetCurrency() string {
	return m.Net.GetCurrency()
}

// LessThan check if this money's gross is less than other's gross
func (t *TaxedMoney) LessThan(other TaxedMoney) bool {
	return t.Gross.LessThan(other.Gross)
}

// Equal checks if two taxed money are equal both in net and gross
func (t *TaxedMoney) Equal(other TaxedMoney) bool {
	return t.Net.Equal(other.Net) && t.Gross.Equal(other.Gross)
}

// LessThanOrEqual checks if this money is less than or equal to other.
func (t *TaxedMoney) LessThanOrEqual(other TaxedMoney) bool {
	return t.LessThan(other) || t.Equal(other)
}

// Mul multiplies current taxed money with given other
//
// other must only be either ints or floats or Decimal
func (m *TaxedMoney) Mul(other float64) TaxedMoney {
	return TaxedMoney{
		Net:   m.Net.Mul(other),
		Gross: m.Gross.Mul(other),
	}
}

// TrueDiv divides current tabled money to other.
// other must be either Decimal or ints or floats
func (t *TaxedMoney) TrueDiv(other float64) TaxedMoney {
	return TaxedMoney{
		Gross: t.Gross.TrueDiv(other),
		Net:   t.Net.TrueDiv(other),
	}
}

// Add adds a money or taxed money to this.
// other must be either Money or TaxedMoney
func (t *TaxedMoney) Add(other any) (*TaxedMoney, error) {
	switch v := other.(type) {
	case Money:
		net, err := t.Net.Add(v)
		if err != nil {
			return nil, err
		}
		gross, err := t.Gross.Add(v)
		if err != nil {
			return nil, err
		}
		return &TaxedMoney{*net, *gross}, nil

	case TaxedMoney:
		net, err := t.Net.Add(v.Net)
		if err != nil {
			return nil, err
		}
		gross, err := t.Gross.Add(v.Gross)
		if err != nil {
			return nil, err
		}
		return &TaxedMoney{*net, *gross}, nil

	default:
		return nil, ErrUnknownType
	}
}

// Add substract this money to other.
// other must be either Money or TaxedMoney.
func (t *TaxedMoney) Sub(other any) (*TaxedMoney, error) {
	switch v := other.(type) {
	case Money:
		net, err := t.Net.Sub(v)
		if err != nil {
			return nil, err
		}
		gross, err := t.Gross.Sub(v)
		if err != nil {
			return nil, err
		}
		return &TaxedMoney{*net, *gross}, nil

	case TaxedMoney:
		net, err := t.Net.Sub(v.Net)
		if err != nil {
			return nil, err
		}
		gross, err := t.Gross.Sub(v.Gross)
		if err != nil {
			return nil, err
		}
		return &TaxedMoney{*net, *gross}, nil

	default:
		return nil, ErrUnknownType
	}
}

// Tax calculates taxed money by subtracting m's gross to m's net
func (t *TaxedMoney) Tax() *Money {
	tax, _ := t.Gross.Sub(t.Net)
	return tax
}

// Return a new instance with both net and gross quantized.
// All arguments are passed to `Money.quantize
func (t *TaxedMoney) Quantize(round Rounding, exp int) (*TaxedMoney, error) {
	net, err := t.Net.Quantize(round, exp)
	if err != nil {
		return nil, err
	}
	gross, err := t.Gross.Quantize(round, exp)
	if err != nil {
		return nil, err
	}

	return &TaxedMoney{
		Net:   *net,
		Gross: *gross,
	}, nil
}

// Apply a fixed discount to TaxedMoney.
func (t *TaxedMoney) fixedDiscount(discount Money) (*TaxedMoney, error) {
	baseNet, err := t.Net.fixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	baseGross, err := t.Gross.fixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	return NewTaxedMoney(*baseNet, *baseGross)
}

func (m *TaxedMoney) fractionalDiscount(fraction decimal.Decimal, fromGross bool) (*TaxedMoney, error) {
	op := Money{
		Currency: m.GetCurrency(),
		Amount:   m.Gross.Amount,
	}
	if !fromGross {
		op.Amount = m.Net.Amount
	}

	op = op.Mul(fraction.InexactFloat64())
	discount, err := op.Quantize(Down, -1)
	if err != nil {
		return nil, err
	}

	return m.fixedDiscount(*discount)
}
