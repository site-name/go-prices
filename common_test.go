package goprices

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestQuantizePrice(t *testing.T) {
	m, err := NewMoney(NewDecimal(decimal.NewFromFloat(23.456)), "vnd")
	if err != nil {
		t.Fatal(err)
	}

	res, err := QuantizePrice(m)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
