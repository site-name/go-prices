package goprices

type TaxedMoneyRange struct {
	start    *TaxedMoney
	stop     *TaxedMoney
	currency string
}

func NewTaxedMoneyRange(start, stop *TaxedMoney) (*TaxedMoneyRange, error) {
	if start.currency != stop.currency {
		return nil, ErrNotSameCurrency
	}
	if less, err := stop.LessThan(start); err == nil && less {
		return nil, ErrStopLessThanStart
	}

	return &TaxedMoneyRange{start, stop, start.currency}, nil
}
