package id

import (
	"crypto/rand"
	"math/big"
)

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type ID string

func NewID(v string) ID {
	return ID(v)
}

func Random() (ID, error) {
	v, err := generate()
	return ID(v), err
}

func MustRandom() ID {
	v, _ := generate()
	return ID(v)
}

func (i ID) String() string {
	return string(i)
}

func generate() (string, error) {
	code := make([]byte, 10)
	for i := range code {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			return "", err
		}
		code[i] = characters[num.Int64()]
	}
	return string(code), nil
}
