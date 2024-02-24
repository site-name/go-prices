package goprices

import (
	"fmt"

	"github.com/site-name/decimal"
)

// Money represents an amount of a particular currency.
type Money struct {
	Amount   decimal.Decimal
	Currency string
}

var (
	_ Currencier             = (*Money)(nil)
	_ MoneyInterface[*Money] = (*Money)(nil)
)

// NewMoney returns new Money object
func NewMoney(amount float64, currency string) (*Money, error) {
	return NewMoneyFromDecimal(decimal.NewFromFloat(amount), currency)
}

func NewMoneyFromDecimal(amount decimal.Decimal, currency string) (*Money, error) {
	unit, err := validateCurrency(currency)
	if err != nil {
		return nil, err
	}
	if amount.LessThan(decimal.Zero) {
		return nil, ErrMoneyNegative
	}
	return &Money{
		Amount:   amount,
		Currency: unit,
	}, nil
}

// String implements fmt.Stringer interface
func (m *Money) String() string {
	return fmt.Sprintf("Money{%s, %s}", m.Amount.String(), m.Currency)
}

// GetCurrency returns current money's Currency
func (m *Money) GetCurrency() string {
	return m.Currency
}

// LessThan checks if other's amount is greater than m's amount
// AND checking same currency included
func (m *Money) LessThan(other Money) bool {
	return m.SameKind(other) && m.Amount.LessThan(other.Amount)
}

// Equal checks if other's amount is equal to m's amount
func (m *Money) Equal(other Money) bool {
	return m.SameKind(other) && m.Amount.Equal(other.Amount)
}

// LessThanOrEqual check if m's amount is less than or equal to other's amount
func (m *Money) LessThanOrEqual(other Money) bool {
	return m.LessThan(other) || m.Equal(other)
}

// Mul multiplty current money with the givent other.
//
// NOTE: other must be either ints or floats or Decimal
func (m *Money) Mul(other float64) Money {
	return Money{
		Currency: m.Currency,
		Amount:   m.Amount.Mul(decimal.NewFromFloat(other)),
	}
}

// TrueDiv divides money with the given other.
//
// NOTE: other must be either ints or uints or floats or Decimal or Money
func (m *Money) TrueDiv(other float64) Money {
	return Money{
		Currency: m.Currency,
		Amount:   m.Amount.DivRound(decimal.NewFromFloat(other), int32(currencies[m.Currency].Fraction)),
	}
}

// Add adds two money amount together, returns new money.
// If returned error is not nil, it could be ErrNotSameCurrency
func (m *Money) Add(other Money) (*Money, error) {
	if !m.SameKind(other) {
		return nil, ErrNotSameCurrency
	}

	return &Money{
		m.Amount.Add(other.Amount),
		m.Currency,
	}, nil
}

// Sub subtracts current money to given other.
// If error is not nil, it could be ErrNotSameCurrency
func (m *Money) Sub(other Money) (*Money, error) {
	if !m.SameKind(other) {
		return nil, ErrNotSameCurrency
	}

	return &Money{
		m.Amount.Sub(other.Amount),
		m.Currency,
	}, nil
}

// func (m *Money) FlatTax(taxRate *decimal.Decimal, kepGross bool) {
// 	faction := decimal.NewFromInt(1).Add(*taxRate)
// 	if kepGross {
// 		// net :=
// 	}
// 	d := decimal.NewFromInt(12)
// }

// Return a copy of the object with its amount quantized.
// NOTE: if exp < 0, default will be used
func (m *Money) Quantize(round Rounding, exp int) (*Money, error) {
	if exp < 0 {
		exp, _ = GetCurrencyPrecision(m.Currency)
	}

	money := &Money{
		Currency: m.Currency,
	}
	switch round {
	case Up:
		money.Amount = m.Amount.RoundUp(int32(exp))
	case Down:
		money.Amount = m.Amount.RoundDown(int32(exp))
	case Ceil:
		money.Amount = m.Amount.RoundCeil(int32(exp))
	case Floor:
		money.Amount = m.Amount.RoundFloor(int32(exp))

	default:
		return nil, ErrInvalidRounding
	}

	return money, nil
}

// Apply a fixed discount to Money type.
func (m *Money) fixedDiscount(discount Money) (*Money, error) {
	sub, err := m.Sub(discount)
	if err != nil {
		return nil, err
	}

	if sub.Amount.GreaterThan(decimal.Zero) {
		return sub, nil
	}

	return &Money{
		Currency: m.Currency,
		Amount:   decimal.Zero,
	}, nil
}

func (m *Money) fractionalDiscount(fraction decimal.Decimal, _ bool) (*Money, error) {
	mul := m.Mul(fraction.InexactFloat64())
	quantized, err := mul.Quantize(Down, -1)
	if err != nil {
		return nil, err
	}

	return m.fixedDiscount(*quantized)
}
