package auth

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var ErrUserNotFound = errors.New("user not found")

type Repository interface {
	FindAll() ([]User, error)
	FindByID(id nanoid.ID) (User, error)
	FindByUsername(u Username) (User, error)
	FindByEmail(e Email) (User, error)
	Save(u User) error
}
