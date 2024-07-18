package goprices

import (
	"fmt"

	"github.com/site-name/decimal"
)

// Money represents an amount of a particular currency.
type Money struct {
	amount   decimal.Decimal
	currency string
}

var _ MoneyInterface[Money] = (*Money)(nil)

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
		amount:   amount,
		currency: unit,
	}, nil
}

// String implements fmt.Stringer interface
func (m Money) String() string {
	return fmt.Sprintf("Money{%s, %s}", m.amount.String(), m.currency)
}

// GetCurrency returns current money's currency
func (m Money) GetCurrency() string {
	return m.currency
}

func (m Money) GetAmount() decimal.Decimal {
	return m.amount
}

// LessThan checks if other's amount is greater than m's amount
// AND checking same currency included
func (m Money) LessThan(other Money) bool {
	return m.SameKind(other) && m.amount.LessThan(other.amount)
}

// Equal checks if other's amount is equal to m's amount
func (m Money) Equal(other Money) bool {
	return m.SameKind(other) && m.amount.Equal(other.amount)
}

// LessThanOrEqual check if m's amount is less than or equal to other's amount
func (m Money) LessThanOrEqual(other Money) bool {
	return m.LessThan(other) || m.Equal(other)
}

// Mul multiplty current money with the givent other.
//
// NOTE: other must be either ints or floats or Decimal
func (m Money) Mul(other float64) Money {
	return Money{
		currency: m.currency,
		amount:   m.amount.Mul(decimal.NewFromFloat(other)),
	}
}

// TrueDiv divides money with the given other.
//
// NOTE: other must be either ints or uints or floats or Decimal or Money
func (m Money) TrueDiv(other float64) Money {
	return Money{
		currency: m.currency,
		amount:   m.amount.DivRound(decimal.NewFromFloat(other), int32(currencies[m.currency].Fraction)),
	}
}

// Add adds two money amount together, returns new money.
// If returned error is not nil, it could be ErrNotSameCurrency
func (m Money) Add(other Money) (*Money, error) {
	if !m.SameKind(other) {
		return nil, ErrNotSameCurrency
	}

	return &Money{
		m.amount.Add(other.amount),
		m.currency,
	}, nil
}

// Neg returns -m
func (m Money) Neg() Money {
	return Money{
		amount:   m.amount.Neg(),
		currency: m.currency,
	}
}

// Sub subtracts current money to given other.
// If error is not nil, it could be ErrNotSameCurrency
func (m Money) Sub(other Money) (*Money, error) {
	return m.Add(m.Neg())
}

// Return a copy of the object with its amount quantized.
// NOTE: if exp < 0, default will be used
func (m Money) Quantize(round Rounding, exp int) (*Money, error) {
	if exp < 0 {
		exp, _ = GetCurrencyPrecision(m.currency)
	}

	money := &Money{
		currency: m.currency,
	}
	switch round {
	case Up:
		money.amount = m.amount.RoundUp(int32(exp))
	case Down:
		money.amount = m.amount.RoundDown(int32(exp))
	case Ceil:
		money.amount = m.amount.RoundCeil(int32(exp))
	case Floor:
		money.amount = m.amount.RoundFloor(int32(exp))

	default:
		return nil, ErrInvalidRounding
	}

	return money, nil
}

// Apply a fixed discount to Money type.
func (m Money) fixedDiscount(discount Money) (*Money, error) {
	sub, err := m.Sub(discount)
	if err != nil {
		return nil, err
	}

	if sub.amount.GreaterThan(decimal.Zero) {
		return sub, nil
	}

	return &Money{
		currency: m.currency,
		amount:   decimal.Zero,
	}, nil
}

func (m Money) fractionalDiscount(fraction decimal.Decimal, _ bool) (*Money, error) {
	mul := m.Mul(fraction.InexactFloat64())
	quantized, err := mul.Quantize(Down, -1)
	if err != nil {
		return nil, err
	}

	return m.fixedDiscount(*quantized)
}
