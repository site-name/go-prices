package goprices

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestTrueDiv(t *testing.T) {
	deci := decimal.NewFromInt(30)
	m, err := NewMoney(&deci, "USD")
	if err != nil {
		t.Fatal(err)
	}
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
	m, err := NewMoney(&deci, "usd")
	if err != nil {
		t.Fatal(err)
	}
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

func TestEqual(t *testing.T) {
	deci := decimal.NewFromInt(20)
	m1, err := NewMoney(&deci, USD)
	if err != nil {
		t.Fatalf("Error NewMoney: %v", err)
	}

	m2 := &Money{&deci, "usd"}

	equal, err := m1.Equal(m2)
	if err != nil {
		t.Fatalf("Error Equal: %v", err)
	}
	if !equal {
		t.Fatal("Error equal result")
	}
	t.Log(m2)
}

func TestQuantize(t *testing.T) {
	deci := decimal.NewFromFloat(20.145)
	m1, err := NewMoney(&deci, USD)
	if err != nil {
		t.Fatalf("Error NewMoney: %v", err)
	}

	fmt.Println(m1)

	m2, err := m1.Quantize()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(m2)
}
