package auth

import "github.com/kindalus/godx/pkg/nanoid"


type User struct {
	ID nanoid.ID
}

func (u User) GetID() nanoid.ID {
	return u.ID
}
