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
	ID          nanoid.ID
	Username    Username
	Email       Email
	PhoneNumber PhoneNumber
	Password    Password
	Role        Role
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
