package crm

import "errors"

var ErrEmptyPhoneNumber = errors.New("phone number is empty")

type PhoneNumber string

func NewPhoneNumber(v string) (PhoneNumber, error) {
	if isEmpty(v) {
		return "", ErrEmptyPhoneNumber
	}

	return PhoneNumber(v), nil
}

func (p PhoneNumber) String() string {
	return string(p)
}
