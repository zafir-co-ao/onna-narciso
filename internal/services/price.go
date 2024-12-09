package services

import "errors"

var ErrInvalidPrice = errors.New("invalid price")

type Price string

func NewPrice(v string) (Price, error) {
	if len(v) == 0 {
		return "", ErrInvalidPrice
	}

	return Price(v), nil
}

func (p Price) Value() string {
	return string(p)
}
