package domain

import errs "github.com/asfqx/SubTrack/internal/errors"

type Currency string

const (
	Rub Currency = "RUB"
	Eur Currency = "EUR"
	Usd Currency = "USD"
)

func NewCurrency(value string) (Currency, error) {
	curr := Currency(value)

	switch curr {
	case Rub, Eur, Usd:
		return curr, nil
	default:
		return "", errs.ErrInvalidCurrency
	}
}
