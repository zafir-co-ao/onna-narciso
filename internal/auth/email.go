package auth

import (
	"errors"
	"regexp"
)

var (
	ErrEmptyEmail         = errors.New("empty email")
	ErrInvalidEmailFormat = errors.New("invalid email format")
)

type Email string

func NewEmail(v string) (Email, error) {
	if len(v) == 0 {
		return Email(v), ErrEmptyEmail
	}

	if !isValidFormat(v) {
		return Email(""), ErrInvalidEmailFormat
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
