package goprices

import (
	"fmt"
	"reflect"

	"github.com/site-name/decimal"
)

// Money represents an amount of a particular currency.
type Money struct {
	Amount   decimal.Decimal
	Currency string
}

var (
	_ Currencyable           = (*Money)(nil)
	_ MoneyInterface[*Money] = (*Money)(nil)
)

// NewMoney returns new Money object
func NewMoney(amount float64, currency string) (*Money, error) {
	unit, err := validateCurrency(currency)
	if err != nil {
		return nil, err
	}
	if amount < 0 {
		return nil, ErrMoneyNegative
	}
	return &Money{
		Amount:   decimal.NewFromFloat(amount),
		Currency: unit,
	}, nil
}

// String implements fmt.Stringer interface
func (m *Money) String() string {
	return fmt.Sprintf("Money{%s, %s}", m.Amount.String(), m.Currency)
}

// MyCurrency returns current money's Currency
func (m *Money) MyCurrency() string {
	return m.Currency
}

// LessThan checks if other's amount is greater than m's amount
// AND checking same currency included
func (m *Money) LessThan(other *Money) bool {
	return m.SameKind(other) && m.Amount.LessThan(other.Amount)
}

// Equal checks if other's amount is equal to m's amount
func (m *Money) Equal(other *Money) bool {
	return m.SameKind(other) && m.Amount.Equal(other.Amount)
}

// LessThanOrEqual check if m's amount is less than or equal to other's amount
func (m *Money) LessThanOrEqual(other *Money) bool {
	return m.LessThan(other) || m.Equal(other)
}

// Mul multiplty current money with the givent other.
//
// NOTE: other must be either ints or floats or Decimal
func (m *Money) Mul(other interface{}) (*Money, error) {
	if other == nil {
		return nil, ErrNillValue
	}

	res := &Money{
		Currency: m.Currency,
	}
	valueOf := reflect.ValueOf(other)

	switch valueOf.Kind() {
	case reflect.Int, reflect.Int8,
		reflect.Int16, reflect.Int32,
		reflect.Int64:
		t := valueOf.Int()
		if t < 0 {
			return nil, ErrMoneyNegative
		}
		res.Amount = m.Amount.Mul(decimal.NewFromInt(t))

	case reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32,
		reflect.Uint64:
		res.Amount = m.Amount.Mul(decimal.NewFromInt(int64(valueOf.Uint())))

	case reflect.Float32, reflect.Float64:
		t := valueOf.Float()
		if t < 0 {
			return nil, ErrMoneyNegative
		}
		res.Amount = m.Amount.Mul(decimal.NewFromFloat(t))

	case reflect.Pointer, reflect.Struct: // for Decimal and *Decimal
		var (
			deci, ok1    = other.(decimal.Decimal)
			deciPtr, ok2 = other.(*decimal.Decimal) // null pointer catched above
		)

		if !(ok1 || ok2) {
			return nil, ErrUnknownType
		}
		if ok2 {
			deci = *deciPtr
		}
		res.Amount = m.Amount.Mul(deci)

	default:
		return nil, ErrUnknownType
	}

	return res, nil
}

// TrueDiv divides money with the given other.
//
// NOTE: other must be either ints or uints or floats or Decimal or Money
func (m *Money) TrueDiv(other interface{}) (*Money, error) {
	if other == nil {
		return nil, ErrNillValue
	}

	res := &Money{
		Currency: m.Currency,
	}
	valueOf := reflect.ValueOf(other)

	switch valueOf.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		t := valueOf.Int()
		if t == 0 {
			return nil, ErrDivisorZero
		} else if t < 0 {
			return nil, ErrMoneyNegative
		}
		res.Amount = m.Amount.Div(decimal.NewFromInt(t))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		t := valueOf.Uint()
		if t == 0 {
			return nil, ErrDivisorZero
		}
		res.Amount = m.Amount.Div(decimal.NewFromInt(int64(t)))

	case reflect.Float32, reflect.Float64:
		t := valueOf.Float()
		if t == 0 {
			return nil, ErrDivisorZero
		} else if t < 0 {
			return nil, ErrMoneyNegative
		}
		res.Amount = m.Amount.Div(decimal.NewFromFloat(t))

	case reflect.Pointer, reflect.Struct:
		var (
			decimalValue decimal.Decimal
			moneyPointer *Money

			mony, ok1    = other.(Money)
			monyPtr, ok2 = other.(*Money) // nil value checked above

			deci, ok3    = other.(decimal.Decimal)
			deciPtr, ok4 = other.(*decimal.Decimal) // nil value checked above
		)

		if ok1 {
			decimalValue = mony.Amount
			moneyPointer = &mony
		} else if ok2 {
			decimalValue = monyPtr.Amount
			moneyPointer = monyPtr
		} else if ok3 {
			decimalValue = deci
		} else if ok4 {
			decimalValue = *deciPtr
		} else { // both are false
			return nil, ErrUnknownType
		}

		if moneyPointer != nil && !m.SameKind(moneyPointer) {
			return nil, ErrNotSameCurrency
		}
		if decimalValue.IsZero() {
			return nil, ErrDivisorZero
		}
		if decimalValue.IsNegative() {
			return nil, ErrMoneyNegative
		}
		res.Amount = res.Amount.Div(decimalValue)

	default:
		return nil, ErrUnknownType
	}

	return res, nil
}

// Add adds two money amount together, returns new money.
// If returned error is not nil, it could be ErrNotSameCurrency
func (m *Money) Add(other any) (*Money, error) {
	if other == nil {
		return nil, ErrNillValue
	}
	mony, ok := other.(*Money)
	if !ok {
		return nil, ErrUnknownType
	}
	if !m.SameKind(mony) {
		return nil, ErrNotSameCurrency
	}
	return &Money{
		m.Amount.Add(mony.Amount),
		m.Currency,
	}, nil
}

// Sub subtracts current money to given other.
// If error is not nil, it could be ErrNotSameCurrency
func (m *Money) Sub(other any) (*Money, error) {
	if other == nil {
		return nil, ErrNillValue
	}
	mony, ok := other.(*Money)
	if !ok {
		return nil, ErrUnknownType
	}
	if !m.SameKind(mony) {
		return nil, ErrNotSameCurrency
	}
	return &Money{
		m.Amount.Sub(mony.Amount),
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
func (m *Money) Quantize(exp *int, round Rounding) (*Money, error) {
	var (
		precision int
		err       error
	)

	if exp != nil && *exp != 0 {
		precision = *exp
	} else {
		precision, err = GetCurrencyPrecision(m.Currency)
		if err != nil {
			return nil, err
		}
	}

	money := &Money{
		Currency: m.Currency,
	}
	switch round {
	case Up:
		money.Amount = m.Amount.RoundUp(int32(precision))
	case Down:
		money.Amount = m.Amount.RoundDown(int32(precision))
	case Ceil:
		money.Amount = m.Amount.RoundCeil(int32(precision))
	case Floor:
		money.Amount = m.Amount.RoundFloor(int32(precision))

	default:
		return nil, ErrInvalidRounding
	}

	return money, nil
}

// Apply a fixed discount to Money type.
func (m *Money) fixedDiscount(discount *Money) (*Money, error) {
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
	mul, err := m.Mul(fraction)
	if err != nil {
		return nil, err
	}
	mul, err = mul.Quantize(nil, Down)
	if err != nil {
		return nil, err
	}

	return m.fixedDiscount(mul)
}
