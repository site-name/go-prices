package goprices

import "errors"

// ErrNotSameCurrency used when manipulate not-same-type money amounts
var (
	ErrNotSameCurrency = errors.New("not same currency")
	ErrUnknownType     = errors.New("unknown given type")
)
