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
	Services      []id.ID
}

func (s *Session) Close(services []string) error {

	if s.IsClosed() {
		return ErrSessionClosed
	}

	for _, v := range services {
		s.Services = append(s.Services, id.NewID(v))
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
