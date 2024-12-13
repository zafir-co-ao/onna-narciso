package auth

type Role string

func (r Role) String() string {
	return string(r)
}
