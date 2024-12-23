package goprices

import (
	"fmt"
	"testing"
)

func TestNewMoneyRange(t *testing.T) {
	start, err := NewMoney(23.45, USD)
	if err != nil {
		t.Fatalf("error create start money: %v", err)
	}
	stop, err := NewMoney(50, USD)
	if err != nil {
		t.Fatalf("error create stop money: %v", err)
	}

	moneyRange, err := NewMoneyRange(*start, *stop)
	if err != nil {
		t.Fatalf("error create new money range: %v", err)
	}

	fmt.Println(moneyRange.String())
}

func TestMoneyRangeSub(t *testing.T) {
	range1, err := NewMoneyRangeFromFloats(23.45, 50, USD)
	if err != nil {
		t.Fatalf("error create range1: %v", err)
	}

	range2, err := NewMoney(2, USD)
	if err != nil {
		t.Fatalf("error create range2: %v", err)
	}

	res, err := range1.Sub(*range2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}
