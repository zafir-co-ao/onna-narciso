package name

import "errors"

var ErrEmptyName = errors.New("the name is empty")

type Name string

func New(v string) (Name, error) {
	if len(v) == 0 {
		return "", ErrEmptyName
	}

	return Name(v), nil
}

func (n Name) String() string {
	return string(n)
}
