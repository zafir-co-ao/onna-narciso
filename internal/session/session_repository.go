package session

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var ErrSessionNotFound = errors.New("session not found")

type Repository interface {
	FindByID(i nanoid.ID) (Session, error)
	Save(s Session) error
}
