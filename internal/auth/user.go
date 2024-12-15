package auth

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var (
	ErrUserNotAllowed     = errors.New("user not allowed")
	ErrOnlyUniqueUsername = errors.New("only unique username")
)

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

func (u *User) IsManager() bool {
	return u.Role == RoleManager
}

func (u *User) VerifyPassword(p string) bool {
	return u.Password.IsValid(p)
}

func (u User) GetID() nanoid.ID {
	return u.ID
}
