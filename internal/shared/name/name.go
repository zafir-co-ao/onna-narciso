package name

import "errors"

var ErrInvalidName = errors.New("the name is empty")

type Name string

func New(v string) (Name, error) {
	if len(v) == 0 {
		return "", ErrInvalidName
	}

	return Name(v), nil
}

func (n Name) String() string {
	return string(n)
}
