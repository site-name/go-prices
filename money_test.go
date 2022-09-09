package goprices

import (
	"fmt"
	"testing"

	"github.com/site-name/decimal"
)

func TestTrueDiv(t *testing.T) {
	m, err := NewMoney(60, "USD")
	if err != nil {
		t.Fatal(err)
	}
	if m == nil {
		t.Fatal("Error creating new money")
	}
	newMoney, err := m.TrueDiv(22.34)
	if err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(newMoney)
	}
}

func TestMul(t *testing.T) {
	m, err := NewMoney(55, "usd")
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
	m1, err := NewMoney(20, USD)
	if err != nil {
		t.Fatalf("Error NewMoney: %v", err)
	}

	m2 := &Money{deci, "usd"}

	equal := m1.Equal(m2)
	if !equal {
		t.Fatal("Error equal result")
	}
	t.Log(m2)
}

func TestQuantize(t *testing.T) {
	m1, err := NewMoney(20.12345, USD)
	if err != nil {
		t.Fatalf("Error NewMoney: %v", err)
	}

	fmt.Println(m1)

	m2, err := m1.Quantize(newInt(4), Up)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(m2)
}

type testCase struct {
	left     Money
	right    Money
	expected bool
}

func TestLessThan(t *testing.T) {

	var testCases = []testCase{
		{
			left: Money{
				Amount:   decimal.NewFromFloat(34.5),
				Currency: USD,
			},
			right: Money{
				Amount:   decimal.NewFromInt(35),
				Currency: USD,
			},
			expected: true,
		},
		{
			left: Money{
				Amount:   decimal.NewFromFloat(34.5),
				Currency: VND,
			},
			right: Money{
				Amount:   decimal.NewFromFloat(79),
				Currency: VND,
			},
			expected: true,
		},
	}

	t.Run("LessThan", func(t *testing.T) {
		for index, testCase := range testCases {

			lessThan := testCase.left.LessThan(&testCase.right)

			if lessThan != testCase.expected {
				t.Fatalf("Case %d: expected: %t, got: %t", index, testCase.expected, lessThan)
			}
		}
	})
}
