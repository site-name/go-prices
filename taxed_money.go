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

func (t *TaxedMoney) GetNet() Money {
	if t == nil {
		panic(ErrNillValue)
	}
	return t.net
}

func (m *TaxedMoney) SetNet(net Money) {
	m.net = net
}

func (m *TaxedMoney) SetGross(gross Money) {
	m.gross = gross
}

func (t *TaxedMoney) GetGross() Money {
	if t == nil {
		panic(ErrNillValue)
	}
	return t.gross
}

// String implements fmt.Stringer interface
func (t *TaxedMoney) String() string {
	if t == nil {
		panic(ErrNillValue)
	}
	return fmt.Sprintf("TaxedMoney{net=%s, gross=%s}", t.net.String(), t.gross.String())
}

// GetCurrency returns current taxed money's currency
func (m *TaxedMoney) GetCurrency() string {
	if m == nil {
		panic(ErrNillValue)
	}
	return m.net.currency
}

// LessThan check if this money's gross is less than other's gross
func (t *TaxedMoney) LessThan(other TaxedMoney) bool {
	return t != nil && t.gross.LessThan(other.gross)
}

// Equal checks if two taxed money are equal both in net and gross
func (t *TaxedMoney) Equal(other TaxedMoney) bool {
	return t != nil && t.net.Equal(other.net) && t.gross.Equal(other.gross)
}

// LessThanOrEqual checks if this money is less than or equal to other.
func (t *TaxedMoney) LessThanOrEqual(other TaxedMoney) bool {
	return t != nil && t.LessThan(other) || t.Equal(other)
}

// Mul multiplies current taxed money with given other
//
// other must only be either ints or floats or Decimal
func (m *TaxedMoney) Mul(other float64) TaxedMoney {
	if m == nil {
		panic(ErrNillValue)
	}
	return TaxedMoney{
		net:   m.net.Mul(other),
		gross: m.gross.Mul(other),
	}
}

// TrueDiv divides current tabled money to other.
// other must be either Decimal or ints or floats
func (t *TaxedMoney) TrueDiv(other float64) TaxedMoney {
	if t == nil {
		panic(ErrNillValue)
	}
	return TaxedMoney{
		gross: t.gross.TrueDiv(other),
		net:   t.net.TrueDiv(other),
	}
}

// Add adds a money or taxed money to this.
// other must be either Money or TaxedMoney
func (t *TaxedMoney) Add(other any) (*TaxedMoney, error) {
	if t == nil || other == nil {
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

func (t *TaxedMoney) Neg() TaxedMoney {
	if t == nil {
		panic(ErrNillValue)
	}
	return TaxedMoney{
		net:   t.net.Neg(),
		gross: t.gross.Neg(),
	}
}

// Add substract this money to other.
// other must be either Money or TaxedMoney.
func (t *TaxedMoney) Sub(other any) (*TaxedMoney, error) {
	if t == nil || other == nil {
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
func (t *TaxedMoney) Tax() *Money {
	if t == nil {
		panic(ErrNillValue)
	}
	tax, _ := t.gross.Sub(t.net)
	return tax
}

// Return a new instance with both net and gross quantized.
// All arguments are passed to `Money.quantize
func (t *TaxedMoney) Quantize(round Rounding, exp int) (*TaxedMoney, error) {
	if t == nil {
		return nil, ErrNillValue
	}
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
func (t *TaxedMoney) fixedDiscount(discount Money) (*TaxedMoney, error) {
	if t == nil {
		return nil, ErrNillValue
	}
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

func (m *TaxedMoney) fractionalDiscount(fraction decimal.Decimal, fromGross bool) (*TaxedMoney, error) {
	if m == nil {
		return nil, ErrNillValue
	}
	op := Money{
		currency: m.GetCurrency(),
		amount:   m.gross.amount,
	}
	if !fromGross {
		op.amount = m.net.amount
	}

	op = op.Mul(fraction.InexactFloat64())
	discount, err := op.Quantize(Down, -1)
	if err != nil {
		return nil, err
	}

	return m.fixedDiscount(*discount)
}
