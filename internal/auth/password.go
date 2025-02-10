package auth

import (
	"crypto/rand"
	"errors"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

var ErrEmptyPassword = errors.New("password is empty")

const PasswordLength = 12

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

func MustNewPassword(v string) Password {
	p, _ := NewPassword(v)
	return p
}

func GeneratePassword(length int) (string, error) {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"
	var password []byte

	for range length {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			return "", err
		}

		password = append(password, characters[n.Int64()])
	}

	return string(password), nil
}

func (p Password) IsValid(v string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.String()), []byte(v))
	return err == nil
}

func (p Password) String() string {
	return string(p)
}
