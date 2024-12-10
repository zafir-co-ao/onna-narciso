package crm

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidEmailFormat = errors.New("Invalid email format")
	ErrEmailAlreadyUsed   = errors.New("Email already used")
)

type Email string

func NewEmail(v string) (Email, error) {
	if len(v) == 0 {
		return Email(""), nil
	}

	if !isValidFormat(v) {
		return "", ErrInvalidEmailFormat
	}
	return Email(v), nil
}

func (e Email) String() string {
	return string(e)
}

func isValidFormat(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
