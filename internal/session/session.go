package session

import (
	"github.com/kindalus/godx/pkg/nanoid"
)

type Session struct {
	ID            nanoid.ID
	AppointmentID nanoid.ID
}

func (s Session) GetID() nanoid.ID {
	return s.ID
}
