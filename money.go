package goprices

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

type Money struct {
	Amount   *decimal.Decimal
	Currency string
}

// ErrNotSameCurrency used when manipulate not-same-type money amounts
var (
	ErrNotSameCurrency = errors.New("not same currency")
	ErrUnknownType     = errors.New("unknown givent type")
)

// NewMoney returns new Money object
func NewMoney(amount *decimal.Decimal, currency string) (*Money, error) {
	code, err := checkCurrency(currency)
	if err != nil {
		return nil, err
	}
	return &Money{
		Amount:   amount,
		Currency: code,
	}, nil
}

// String implements fmt.Stringer interface
func (m *Money) String() string {
	return fmt.Sprintf("Money{%q, %q}", m.Amount.String(), m.Currency)
}

func (m *Money) sameKind(other *Money) error {
	if m.Currency == other.Currency {
		return nil
	}
	return ErrNotSameCurrency
}

func (m *Money) LessThan(other *Money) (bool, error) {
	err := m.sameKind(other)
	if err != nil {
		return false, err
	}
	return m.Amount.LessThan(*other.Amount), nil
}

func (m *Money) Equal(other *Money) (bool, error) {
	err := m.sameKind(other)
	if err != nil {
		return false, err
	}
	return m.Amount.Equal(*other.Amount), nil
}

// LessThanOrEqual check if this money is less than or equal to other
func (m *Money) LessThanOrEqual(other *Money) (bool, error) {
	less, err := m.LessThan(other)
	if err != nil {
		return false, err
	}
	eq, err := m.Equal(other)
	if err != nil {
		return false, err
	}
	return less || eq, nil
}

// Mul multiplty money with the givent other.
// other must be a *Money, float64, float32, int64, int, int32
func (m *Money) Mul(other interface{}) (*Money, error) {
	switch t := other.(type) {
	case *Money:
		d := m.Amount.Mul(*t.Amount)
		return NewMoney(&d, m.Currency)
	case float64:
		floatDeci := decimal.NewFromFloat(t)
		d := m.Amount.Mul(floatDeci)
		return NewMoney(&d, m.Currency)
	case float32:
		floatDeci := decimal.NewFromFloat32(t)
		d := m.Amount.Mul(floatDeci)
		return NewMoney(&d, m.Currency)
	case int64:
		intDeci := decimal.NewFromInt(t)
		d := m.Amount.Mul(intDeci)
		return NewMoney(&d, m.Currency)
	case int:
		intDeci := decimal.NewFromInt32(int32(t))
		d := m.Amount.Mul(intDeci)
		return NewMoney(&d, m.Currency)
	case int32:
		intDeci := decimal.NewFromInt32(t)
		d := m.Amount.Mul(intDeci)
		return NewMoney(&d, m.Currency)

	default:
		return nil, ErrUnknownType
	}
}

// Mul divide money with the givent other.
// other must be a *Money, float64, float32, int64, int, int32
func (m *Money) TrueDiv(other interface{}) (*Money, error) {
	switch t := other.(type) {
	case *Money:
		if err := m.sameKind(t); err != nil {
			return nil, err
		}
		d := m.Amount.Div(*t.Amount)
		return NewMoney(&d, m.Currency)
	case float64:
		floatDeci := decimal.NewFromFloat(t)
		d := m.Amount.Div(floatDeci)
		return NewMoney(&d, m.Currency)
	case float32:
		floatDeci := decimal.NewFromFloat32(t)
		d := m.Amount.Div(floatDeci)
		return NewMoney(&d, m.Currency)
	case int64:
		intDeci := decimal.NewFromInt(t)
		d := m.Amount.Div(intDeci)
		return NewMoney(&d, m.Currency)
	case int:
		intDeci := decimal.NewFromInt32(int32(t))
		d := m.Amount.Div(intDeci)
		return NewMoney(&d, m.Currency)
	case int32:
		intDeci := decimal.NewFromInt32(t)
		d := m.Amount.Div(intDeci)
		return NewMoney(&d, m.Currency)

	default:
		return nil, ErrUnknownType
	}
}

// Add adds two money amount together, returns new instance of money
func (m *Money) Add(other *Money) (*Money, error) {
	if err := m.sameKind(other); err != nil {
		return nil, err
	}
	amount := m.Amount.Add(*other.Amount)
	return &Money{&amount, m.Currency}, nil
}

// Sub reduce two money amount and returns new instance of money
func (m *Money) Sub(other *Money) (*Money, error) {
	if err := m.sameKind(other); err != nil {
		return nil, err
	}
	amount := m.Amount.Sub(*other.Amount)
	return &Money{&amount, m.Currency}, nil
}

func (m *Money) IsNotZero() bool {
	return !m.Amount.IsZero()
}

// Return a copy of the object with its amount quantized.
// If `exp` is given the resulting exponent will match that of `exp`.
// Otherwise the resulting exponent will be set to the correct exponent
// of the currency if it's known and to default (two decimal places)
// otherwise.
// func (m *Money) Quantize(exp *decimal.Decimal) *Money {
// 	if exp == nil {

// 	}
// }
