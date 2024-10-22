package scheduling

import (
	"crypto/rand"
	"math/big"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type ID string

func NewID(v string) ID {
	return ID(v)
}

func Random() (ID, error) {
	v, err := generate()
	return ID(v), err
}

func (i ID) Value() string {
	return string(i)
}

func generate() (string, error) {
	code := make([]byte, 10)
	for i := range code {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		code[i] = letters[num.Int64()]
	}
	return string(code), nil
}
