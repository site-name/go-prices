package main

import (
	"fmt"

	"github.com/site-name/decimal"
)

func main() {
	d := decimal.NewFromFloat(125.456)
	fmt.Println(d.Round(-3))
}
