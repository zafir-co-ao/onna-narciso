package crm

import (
	"errors"
	"regexp"
)

var ErrInvalidFormat = errors.New("invalid email format")

type Email string

func NewEmail(v string) (Email, error) {
	if !isValidFormat(v) {
		return "", ErrInvalidFormat
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
