package goprices

import (
	"fmt"
	"testing"

	"github.com/site-name/decimal"
)

func Test_FixedDiscount(t *testing.T) {
	discount, err := NewMoney(23.45, "JPY")
	if err != nil {
		t.Fatal(err)
	}
	m, err := NewMoney(45, "JPY")
	if err != nil {
		t.Fatal(err)
	}
	value, err := FixedDiscount[*Money](m, discount)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(value)
}

func Test_FractionalDiscount(t *testing.T) {
	m, err := NewMoney(100.456, "USD")
	if err != nil {
		t.Fatal(err)
	}

	m, err = FractionalDiscount[*Money](m, decimal.NewFromFloat(0.2), false)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(m)

	mRange, err := NewMoneyRange(
		&Money{
			decimal.NewFromFloat(400.67),
			"VND",
		},
		&Money{
			decimal.NewFromFloat(800.2365),
			"VND",
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	iface, err := FractionalDiscount[*MoneyRange](mRange, decimal.NewFromFloat(0.135), true)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(iface)
}

func Test_PercentageDiscount(t *testing.T) {
	m := &Money{
		Amount:   decimal.NewFromFloat(566.63),
		Currency: "usd",
	}
	vl, err := PercentageDiscount[*Money](m, 50.0, true)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(vl)
}
