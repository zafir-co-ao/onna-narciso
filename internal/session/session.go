package session

import "github.com/zafir-co-ao/onna-narciso/internal/shared/id"

type Session struct {
	ID            id.ID
	AppointmentID id.ID
}

func (s Session) GetID() id.ID {
	return s.ID
}
