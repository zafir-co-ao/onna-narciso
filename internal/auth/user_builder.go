package auth

import (
	"github.com/kindalus/godx/pkg/nanoid"
)

type userBuilder User

func NewUserBuilder() *userBuilder {
	return &userBuilder{ID: nanoid.New()}
}

func (u *userBuilder) WithID(id nanoid.ID) *userBuilder {
	u.ID = id
	return u
}

func (u *userBuilder) WithUserName(n Username) *userBuilder {
	u.Username = n
	return u
}

func (u *userBuilder) WithEmail(e Email) *userBuilder {
	u.Email = e
	return u
}

func (u *userBuilder) WithPhoneNumber(p PhoneNumber) *userBuilder {
	u.PhoneNumber = p
	return u
}

func (u *userBuilder) WithPassWord(p Password) *userBuilder {
	u.Password = p
	return u
}

func (u *userBuilder) WithRole(r Role) *userBuilder {
	u.Role = r
	return u
}

func (u *userBuilder) Build() User {
	return User{
		ID:          u.ID,
		Username:    u.Username,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
		Password:    u.Password,
		Role:        u.Role,
	}
}
