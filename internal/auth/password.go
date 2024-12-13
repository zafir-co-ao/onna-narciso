package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrEmptyPassword = errors.New("password is empty")

type Password string

func NewPassword(v string) (Password, error) {
	if len(v) == 0 {
		return Password(""), ErrEmptyPassword
	}

	b, err := bcrypt.GenerateFromPassword([]byte(v), bcrypt.MinCost)
	if err != nil {
		return Password(""), err
	}

	return Password(string(b)), nil
}

func (p Password) IsValid(v string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.String()), []byte(v))
	return err == nil
}

func (p Password) String() string {
	return string(p)
}
