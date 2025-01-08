package sessions

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
)

var ErrSessionNotFound = errors.New("session not found")

type Repository interface {
	FindByID(id nanoid.ID) (Session, error)
	FindByAppointmentsIDs(ids []nanoid.ID) ([]Session, error)
	Save(s Session) error
}
