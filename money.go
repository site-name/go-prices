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

// NewMoney returns new Money object
func NewMoney(amount float64, currency string) (*Money, error) {
	unit, err := checkCurrency(currency)
	if err != nil {
		return nil, err
	}
	return &Money{
		Amount:   decimal.NewFromFloat(amount),
		Currency: unit,
	}, nil
}

// String implements fmt.Stringer interface
func (m *Money) String() string {
	return fmt.Sprintf("Money{%q, %q}", m.Amount.String(), m.Currency)
}

// MyCurrency returns current money's Currency
func (m *Money) MyCurrency() string {
	return m.Currency
}

// LessThan checks if other's amount is greater than m's amount
// AND checking same currency included
func (m *Money) LessThan(other *Money) (bool, error) {
	if err := m.SameKind(other); err != nil {
		return false, err
	}
	return m.Amount.LessThan(other.Amount), nil
}

// Equal checks if other's amount is equal to m's amount
func (m *Money) Equal(other *Money) (bool, error) {
	if err := m.SameKind(other); err != nil {
		return false, err
	}
	return m.Amount.Equal(other.Amount), nil
}

// LessThanOrEqual check if m's amount is less than or equal to other's amount
func (m *Money) LessThanOrEqual(other *Money) (bool, error) {
	if err := m.SameKind(other); err != nil {
		return false, err
	}

	return m.Amount.LessThanOrEqual(other.Amount), nil
}

// Mul multiplty money with the givent other.
//
// NOTE: other must be either `int` or `float64` or `Decimal`
func (m *Money) Mul(other interface{}) (*Money, error) {
	switch t := other.(type) {
	case int:
		return &Money{
			Amount:   m.Amount.Mul(decimal.NewFromInt32(int32(t))),
			Currency: m.Currency,
		}, nil

	case float64:
		return &Money{
			Amount:   m.Amount.Mul(decimal.NewFromFloat(t)),
			Currency: m.Currency,
		}, nil

	case decimal.Decimal:
		return &Money{
			Amount:   m.Amount.Mul(t),
			Currency: m.Currency,
		}, nil

	default:
		return nil, ErrUnknownType
	}
}

// TrueDiv divides money with the given other.
//
// NOTE: other must be either `*Money` or `int` or `float64` or `Decimal`
func (m *Money) TrueDiv(other interface{}) (*Money, error) {
	res := &Money{
		Currency: m.Currency,
	}

	switch t := other.(type) {
	case int:
		if t == 0 {
			return nil, ErrDivisorNotZero
		}
		res.Amount = m.Amount.Div(decimal.NewFromInt32(int32(t)))

	case float64:
		res.Amount = m.Amount.Div(decimal.NewFromFloat(t))

	case *Money:
		if t.Amount.IsZero() {
			return nil, ErrDivisorNotZero
		}
		if err := m.SameKind(t); err != nil {
			return nil, err
		}
		res.Amount = m.Amount.Div(t.Amount)

	case decimal.Decimal:
		if t.IsZero() {
			return nil, ErrDivisorNotZero
		}
		res.Amount = m.Amount.Div(t)

	default:
		return nil, ErrUnknownType
	}

	return res, nil
}

// Add adds two money amount together, returns new money
func (m *Money) Add(other *Money) (*Money, error) {
	if err := m.SameKind(other); err != nil {
		return nil, err
	}
	return &Money{
		m.Amount.Add(other.Amount),
		m.Currency,
	}, nil
}

// Sub subtracts currenct money to given `money`
func (m *Money) Sub(other *Money) (*Money, error) {
	if err := m.SameKind(other); err != nil {
		return nil, err
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
// If `exp` is given the resulting exponent will match that of `exp`.
// Otherwise the resulting exponent will be set to the correct exponent
// of the currency if it's known and to default (two decimal places)
// otherwise.
func (m *Money) Quantize(round Rounding) (*Money, error) {
	precision, err := GetCurrencyPrecision(m.Currency)
	if err != nil {
		return nil, err
	}

	var roundFunc RoundFunc = nil

	switch round {
	case Up:
		roundFunc = m.Amount.RoundUp
	case Down:
		roundFunc = m.Amount.RoundDown
	case Ceil:
		roundFunc = m.Amount.RoundCeil
	case Floor:
		roundFunc = m.Amount.RoundFloor

	default:
		return nil, ErrInvalidRounding
	}

	return &Money{
		Amount:   roundFunc(precision),
		Currency: m.Currency,
	}, nil
}

// Apply a fixed discount to Money type.
func (m *Money) FixedDiscount(discount *Money) (*Money, error) {
	sub, err := m.Sub(discount) // same currencies check included
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
