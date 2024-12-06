package phone

import "errors"

var ErrEmptyPhoneNumber = errors.New("phone number is empty")

type PhoneNumber string

func New(v string) (PhoneNumber, error) {
	if isEmpty(v) {
		return "", ErrEmptyPhoneNumber
	}

	return PhoneNumber(v), nil
}

func (p PhoneNumber) String() string {
	return string(p)
}

func isEmpty(v string) bool {
	return len(v) == 0
}
