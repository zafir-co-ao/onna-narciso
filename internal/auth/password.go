package auth

type Password string

func (p Password) String() string {
	return string(p)
}
