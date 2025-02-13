package auth

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var (
	ErrUserNotAllowed        = errors.New("user not allowed")
	ErrOnlyUniqueUsername    = errors.New("only unique username")
	ErrOnlyUniqueEmail       = errors.New("only unique email")
	ErrOnlyUniquePhoneNumber = errors.New("only unique phone number")
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

func (u *User) VerifyPassword(v string) bool {
	return u.Password.IsValid(v)
}

func (u *User) IsPasswordConfirmed(newPwd, confirmPwd string) bool {
	return newPwd == confirmPwd
}

func (u *User) UpdatePassword(p Password) {
	u.Password = p
}

func (u *User) ResetPassword(v string) {
	p := MustNewPassword(v)
	u.UpdatePassword(p)
}

func (u User) GetID() nanoid.ID {
	return u.ID
}
