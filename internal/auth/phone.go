package auth

import "errors"

type PhoneNumber string

var ErrEmptyPhoneNumber = errors.New("empty phone number")

func NewPhoneNumber(v string) (PhoneNumber, error) {
	if len(v) == 0 {
		return PhoneNumber(v), ErrEmptyPhoneNumber
	}

	return PhoneNumber(v), nil
}

func (p PhoneNumber) String() string {
	return string(p)
}
