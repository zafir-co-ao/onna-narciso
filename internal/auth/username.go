package auth

import "errors"

type Username string

var ErrEmptyUsername = errors.New("user name is empty")

func NewUsername(v string) (Username, error) {
	if len(v) == 0 {
		return Username(""), ErrEmptyUsername
	}

	return Username(v), nil
}

func (u Username) String() string {
	return string(u)
}
