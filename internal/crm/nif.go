package crm

import "errors"

var ErrEmptyNif = errors.New("nif is empty")

type Nif string

func NewNif(v string) (Nif, error) {
	if len(v) == 0 {
		return "", ErrEmptyNif
	}

	return Nif(v), nil
}

func (n Nif) String() string {
	return string(n)
}
