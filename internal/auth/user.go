package auth

import "github.com/kindalus/godx/pkg/nanoid"

type User struct {
	ID       nanoid.ID
	Username Username
	Password Password
	Role     Role
}

func (u User) GetID() nanoid.ID {
	return u.ID
}
