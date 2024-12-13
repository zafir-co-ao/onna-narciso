package auth

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)


var ErrUserNotFound = errors.New("user not found")

type Repository interface {
	FindByID(id nanoid.ID) (User, error)
	Save(u User) (error)
}
