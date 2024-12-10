package crm

import "errors"

var ErrPhoneNumberAlreadyUsed = errors.New("Phone number already used")

type PhoneNumber string

func NewPhoneNumber(v string) (PhoneNumber, error) {
	return PhoneNumber(v), nil
}

func (p PhoneNumber) String() string {
	return string(p)
}
