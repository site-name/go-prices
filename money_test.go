package goprices

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestTrueDiv(t *testing.T) {
	deci := decimal.NewFromInt(30)
	m, err := NewMoney(&deci, "USD")
	if m == nil {
		t.Fatal("Error creating new money")
	}
	newMoney, err := m.TrueDiv(22.1212)
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(newMoney)
	}
}

func TestMul(t *testing.T) {
	deci := decimal.NewFromInt(30)
	m, err := NewMoney(&deci, "USD")
	if m == nil {
		t.Fatal("Error creating new money")
	}
	newMoney, err := m.Mul(22.34)
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(newMoney)
	}
}
