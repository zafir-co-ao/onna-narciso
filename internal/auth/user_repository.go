package auth

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrAutenticateFailed = errors.New("credentials invalid")
)

type Repository interface {
	FindByID(id nanoid.ID) (User, error)
	FindByUserName(un Username) (User, error)
	Save(u User) error
}
