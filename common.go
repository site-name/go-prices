package goprices

// QuantizePrice accepts the `price` argument to be either:
// *TaxedMoney or *Money or *TaxedMoneyRange
func QuantizePrice(price interface{}) (interface{}, error) {
	switch v := price.(type) {
	case *TaxedMoney:
		return v.Quantize()
	case *Money:
		return v.Quantize()
	case *TaxedMoneyRange:
		return v.Quantize()

	default:
		return nil, ErrUnknownType
	}
}
