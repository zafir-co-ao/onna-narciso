package auth

import "github.com/kindalus/godx/pkg/nanoid"

type User struct {
	ID       nanoid.ID
	Username Username
	Password Password
	Role     Role
}

func NewUser(u Username, p Password, r Role) User {
	return User{
		ID:       nanoid.New(),
		Username: u,
		Password: p,
		Role:     r,
	}
}

func (u *User) VerifyPassword(p string) bool {
	return u.Password.IsValid(p)
}

func (u User) GetID() nanoid.ID {
	return u.ID
}

func (u User) IsSamePassword(p Password) bool {
	return u.Password == p
}
