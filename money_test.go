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
	newMoney := m.TrueDiv(22.34)
	fmt.Println(newMoney)
}

func TestMul(t *testing.T) {
	m, err := NewMoney(55, "usd")
	if err != nil {
		t.Fatal(err)
	}
	if m == nil {
		t.Fatal("Error creating new money")
	}
	newMoney := m.Mul(22.34)
	fmt.Println(newMoney)
}

func TestEqual(t *testing.T) {
	deci := decimal.NewFromInt(20)
	m1, err := NewMoney(20.0, USD)
	if err != nil {
		t.Fatalf("Error NewMoney: %v", err)
	}

	m2 := Money{deci, USD}

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

	m2, err := m1.Quantize(Up, -1)
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
				amount:   decimal.NewFromFloat(34.5),
				currency: USD,
			},
			right: Money{
				amount:   decimal.NewFromInt(35),
				currency: USD,
			},
			expected: true,
		},
		{
			left: Money{
				amount:   decimal.NewFromFloat(34.5),
				currency: VND,
			},
			right: Money{
				amount:   decimal.NewFromFloat(79),
				currency: VND,
			},
			expected: true,
		},
	}

	t.Run("LessThan", func(t *testing.T) {
		for index, testCase := range testCases {

			lessThan := testCase.left.LessThan(testCase.right)

			if lessThan != testCase.expected {
				t.Fatalf("Case %d: expected: %t, got: %t", index, testCase.expected, lessThan)
			}
		}
	})
}

func TestSub(t *testing.T) {
	m1, err := NewMoney(45, USD)
	if err != nil {
		t.Fatal(err)
	}

	m2, err := NewMoney(23, USD)
	if err != nil {
		t.Fatal(err)
	}

	sub, err := m1.Sub(*m2)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(sub)
}

func TestAdd(t *testing.T) {
	m1, err := NewMoney(45, USD)
	if err != nil {
		t.Fatal(err)
	}

	m2, err := NewMoney(23, USD)
	if err != nil {
		t.Fatal(err)
	}

	add, err := m1.Add(*m2)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(add)
}

func TestNeg(t *testing.T) {
	m1, err := NewMoney(45, USD)
	if err != nil {
		t.Fatal(err)
	}

	neg := m1.Neg()

	fmt.Println(neg)
}
