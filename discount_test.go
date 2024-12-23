package goprices

import (
	"fmt"
	"testing"

	"github.com/site-name/decimal"
)

func TestFixedDiscount(t *testing.T) {
	discount, err := NewMoney(23.45, JPY)
	if err != nil {
		t.Fatal(err)
	}
	m, err := NewMoney(45, JPY)
	if err != nil {
		t.Fatal(err)
	}
	value, err := FixedDiscount(m, *discount)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(value)
}

func TestFractionalDiscount(t *testing.T) {
	m, err := NewMoney(100.456, "USD")
	if err != nil {
		t.Fatal(err)
	}

	m, err = FractionalDiscount(m, decimal.NewFromFloat(0.2), false, Up)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(m)

	mRange, err := NewMoneyRangeFromFloats(400.67, 800.2365, VND)
	if err != nil {
		t.Fatal(err)
	}

	iface, err := FractionalDiscount(mRange, decimal.NewFromFloat(0.135), true, Up)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(iface)
}

func TestPercentageDiscount(t *testing.T) {
	m := &Money{
		amount:   decimal.NewFromFloat(566.63),
		currency: "usd",
	}
	vl, err := PercentageDiscount(m, 50.0, true, Up)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(vl)
}
