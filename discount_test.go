package goprices

import (
	"fmt"
	"testing"

	"github.com/site-name/decimal"
)

func Test_FixedDiscount(t *testing.T) {

	discount, err := NewMoney(decimal.NewFromFloat(23.45), "JPY")
	if err != nil {
		t.Fatal(err)
	}
	m, err := NewMoney(decimal.NewFromInt(45), "JPY")
	if err != nil {
		t.Fatal(err)
	}
	value, err := FixedDiscount(m, discount)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(value.(*Money))
}

func Test_FractionalDiscount(t *testing.T) {
	mRange, err := NewMoneyRange(
		&Money{
			decimal.NewFromFloat(34.67),
			"VND",
		},
		&Money{
			decimal.NewFromFloat(800.2365),
			"VND",
		},
	)
	fmt.Println(mRange)
	if err != nil {
		t.Fatal(err)
	}
	iface, err := FractionalDiscount(mRange, decimal.NewFromFloat(13.5), true)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(iface.(*MoneyRange))
}

func Test_PercentageDiscount(t *testing.T) {
	m := &Money{
		Amount:   decimal.NewFromFloat(566.64),
		Currency: "usd",
	}
	vl, err := PercentageDiscount(m, 24, true)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(vl.(*Money))
}
