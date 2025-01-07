package services

import (
	"errors"
	"strconv"
)

const (
	MinDiscount = 0
	MaxDiscount = 100
)

var ErrDiscountNotAllowed = errors.New("discount must be between 0 and 100 percentage")

type Discount string

func NewDiscount(v string) (Discount, error) {
	if v == "" {
		return Discount(v), nil
	}

	if !isValid(v) {
		return Discount(""), ErrDiscountNotAllowed
	}

	return Discount(v), nil
}

func isValid(v string) bool {
	va, err := strconv.Atoi(v)

	if err != nil {
		return false
	}

	return va <= MaxDiscount && va >= MinDiscount
}

func (d Discount) String() string {
	return string(d)
}
