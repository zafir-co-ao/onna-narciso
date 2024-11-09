package session

import (
	"time"

	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

const StatusClosed Status = "closed"

type Status string

type Session struct {
	ID            id.ID
	AppointmentID id.ID
	Status        Status
	CloseTime     time.Time
}

func (s *Session) Close() {
	s.CloseTime = time.Now()
	s.Status = StatusClosed
}

func (s *Session) IsClosed() bool {
	return s.Status == StatusClosed
}

func (s Session) GetID() id.ID {
	return s.ID
}
