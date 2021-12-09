package goprices

import (
	"fmt"
	"testing"
)

func TestNewMoneyRange(t *testing.T) {
	start, err := NewMoney(2345, USD)
	if err != nil {
		t.Fatalf("error create start money: %v", err)
	}
	stop, err := NewMoney(50, USD)
	if err != nil {
		t.Fatalf("error create stop money: %v", err)
	}

	moneyRange, err := NewMoneyRange(start, stop)
	if err != nil {
		t.Fatalf("error create new money range: %v", err)
	}

	fmt.Println(moneyRange.String())
}
