package price

import "errors"

var ErrInvalidPrice = errors.New("invalid price")

type Price string

func New(v string) (Price, error) {
	if len(v) == 0 {
		return "", ErrInvalidPrice
	}

	return Price(v), nil
}
