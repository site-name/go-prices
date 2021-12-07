package goprices

import (
	"fmt"
	"testing"
)

func TestQuantizePrice(t *testing.T) {
	m, err := NewMoney(23.456, "vnd")
	if err != nil {
		t.Fatal(err)
	}

	res, err := QuantizePrice(m)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(res)
}
