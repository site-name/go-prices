package goprices

import (
	"fmt"

	"github.com/site-name/decimal"
)

// TaxedMoney represents taxed money. It wraps net, gross money and currency.
type TaxedMoney struct {
	net   Money
	gross Money
}

var _ MoneyInterface[TaxedMoney] = (*TaxedMoney)(nil)

// NewTaxedMoney returns new TaxedMoney,
// If net and gross have different currency type, return nil and error
func NewTaxedMoney(net, gross Money) (*TaxedMoney, error) {
	unit1, err := validateCurrency(net.currency)
	if err != nil {
		return nil, err
	}
	unit2, err := validateCurrency(gross.currency)
	if err != nil {
		return nil, err
	}

	if unit1 != unit2 {
		return nil, ErrNotSameCurrency
	}

	return &TaxedMoney{net, gross}, nil
}

func NewTaxedMoneyFromFloats(net, gross float64, currency string) (*TaxedMoney, error) {
	netMoney, err := NewMoney(net, currency)
	if err != nil {
		return nil, err
	}
	grossMoney, err := NewMoney(gross, currency)
	if err != nil {
		return nil, err
	}
	return NewTaxedMoney(*netMoney, *grossMoney)
}

func NewTaxedMoneyFromDecimals(net, gross decimal.Decimal, currency string) (*TaxedMoney, error) {
	netMoney, err := NewMoneyFromDecimal(net, currency)
	if err != nil {
		return nil, err
	}
	grossMoney, err := NewMoneyFromDecimal(gross, currency)
	if err != nil {
		return nil, err
	}
	return NewTaxedMoney(*netMoney, *grossMoney)
}

func (t TaxedMoney) GetNet() Money {
	return t.net
}

func (m *TaxedMoney) SetNet(net Money) {
	m.net = net
}

func (m *TaxedMoney) SetGross(gross Money) {
	m.gross = gross
}

func (t TaxedMoney) GetGross() Money {
	return t.gross
}

// String implements fmt.Stringer interface
func (t TaxedMoney) String() string {
	return fmt.Sprintf("TaxedMoney{net=%s, gross=%s}", t.net.String(), t.gross.String())
}

// GetCurrency returns current taxed money's currency
func (m TaxedMoney) GetCurrency() string {
	return m.net.currency
}

// LessThan check if this money's gross is less than other's gross
func (t TaxedMoney) LessThan(other TaxedMoney) bool {
	return t.gross.LessThan(other.gross)
}

// Equal checks if two taxed money are equal both in net and gross
func (t TaxedMoney) Equal(other TaxedMoney) bool {
	return t.net.Equal(other.net) && t.gross.Equal(other.gross)
}

// LessThanOrEqual checks if this money is less than or equal to other.
func (t TaxedMoney) LessThanOrEqual(other TaxedMoney) bool {
	return t.LessThan(other) || t.Equal(other)
}

// Mul multiplies current taxed money with given other
//
// other must only be either ints or floats or Decimal
func (t TaxedMoney) Mul(other float64) TaxedMoney {
	return TaxedMoney{
		net:   t.net.Mul(other),
		gross: t.gross.Mul(other),
	}
}

// TrueDiv divides current tabled money to other.
// other must be either Decimal or ints or floats
func (t TaxedMoney) TrueDiv(other float64) TaxedMoney {
	return TaxedMoney{
		gross: t.gross.TrueDiv(other),
		net:   t.net.TrueDiv(other),
	}
}

// Add adds a money or taxed money to this.
// other must be either Money or TaxedMoney
func (t TaxedMoney) Add(other any) (*TaxedMoney, error) {
	if other == nil {
		return nil, ErrNillValue
	}

	switch v := other.(type) {
	case Money:
		net, err := t.net.Add(v)
		if err != nil {
			return nil, err
		}
		gross, err := t.gross.Add(v)
		if err != nil {
			return nil, err
		}
		return &TaxedMoney{*net, *gross}, nil

	case TaxedMoney:
		net, err := t.net.Add(v.net)
		if err != nil {
			return nil, err
		}
		gross, err := t.gross.Add(v.gross)
		if err != nil {
			return nil, err
		}
		return &TaxedMoney{*net, *gross}, nil

	default:
		return nil, ErrUnknownType
	}
}

func (t TaxedMoney) Neg() TaxedMoney {
	return TaxedMoney{
		net:   t.net.Neg(),
		gross: t.gross.Neg(),
	}
}

// Add substract this money to other.
// other must be either Money or TaxedMoney.
func (t TaxedMoney) Sub(other any) (*TaxedMoney, error) {
	if other == nil {
		return nil, ErrNillValue
	}

	switch v := other.(type) {
	case Money:
		return t.Add(v.Neg())
	case TaxedMoney:
		return t.Add(v.Neg())

	default:
		return nil, ErrUnknownType
	}
}

// Tax calculates taxed money by subtracting m's gross to m's net
func (t TaxedMoney) Tax() *Money {
	tax, _ := t.gross.Sub(t.net)
	return tax
}

// Return a new instance with both net and gross quantized.
// All arguments are passed to `Money.quantize
func (t TaxedMoney) Quantize(round Rounding, exp int) (*TaxedMoney, error) {
	net, err := t.net.Quantize(round, exp)
	if err != nil {
		return nil, err
	}
	gross, err := t.gross.Quantize(round, exp)
	if err != nil {
		return nil, err
	}

	return &TaxedMoney{
		net:   *net,
		gross: *gross,
	}, nil
}

// Apply a fixed discount to TaxedMoney.
func (t TaxedMoney) fixedDiscount(discount Money) (*TaxedMoney, error) {
	baseNet, err := t.net.fixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	baseGross, err := t.gross.fixedDiscount(discount)
	if err != nil {
		return nil, err
	}
	return NewTaxedMoney(*baseNet, *baseGross)
}

func (m TaxedMoney) fractionalDiscount(fraction decimal.Decimal, fromGross bool, rounding Rounding) (*TaxedMoney, error) {
	op := Money{
		currency: m.GetCurrency(),
		amount:   m.gross.amount,
	}
	if !fromGross {
		op.amount = m.net.amount
	}

	op = op.Mul(fraction.InexactFloat64())
	discount, err := op.Quantize(rounding, -1)
	if err != nil {
		return nil, err
	}

	return m.fixedDiscount(*discount)
}
