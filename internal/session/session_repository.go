package session

import (
	"errors"

	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

var ErrSessionNotFound = errors.New("session not found")

type SessionRepository interface {
	FindByID(i id.ID) (Session, error)
	Save(s Session) error
}
