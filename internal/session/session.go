package session

import (
	"errors"
	"time"

	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

const StatusClosed Status = "closed"

var ErrSessionClosed = errors.New("Session already closed")

type Status string

type Session struct {
	ID            id.ID
	AppointmentID id.ID
	Status        Status
	CloseTime     time.Time
}

func (s *Session) Close() error {

	if s.IsClosed() {
		return ErrSessionClosed
	}

	s.CloseTime = time.Now()
	s.Status = StatusClosed
	return nil
}

func (s *Session) IsClosed() bool {
	return s.Status == StatusClosed
}

func (s Session) GetID() id.ID {
	return s.ID
}
