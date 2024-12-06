package nif

import "errors"

var ErrEmptyNif = errors.New("nif is empty")

type Nif string

func New(v string) (Nif, error) {
	if isEmpty(v) {
		return "", ErrEmptyNif
	}

	return Nif(v), nil
}

func (n Nif) String() string {
	return string(n)
}

func isEmpty(v string) bool {
	return len(v) == 0
}
