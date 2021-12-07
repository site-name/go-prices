package goprices

// QuantizePrice accepts the `price` argument to be either:
// *TaxedMoney or *Money or *TaxedMoneyRange
func QuantizePrice(price interface{}, round Rounding) (interface{}, error) {
	switch v := price.(type) {
	case *TaxedMoney:
		return v.Quantize(round)
	case *Money:
		return v.Quantize(round)
	case *TaxedMoneyRange:
		return v.Quantize(round)

	default:
		return nil, ErrUnknownType
	}
}
