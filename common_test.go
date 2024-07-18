package goprices

import (
	"fmt"
	"testing"
)

func TestQuantizePrice(t *testing.T) {
	money, err := NewMoney(34.497, "USD")
	if err != nil {
		t.Fatal(err)
	}

	res, err := QuantizePrice(money, Up)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
