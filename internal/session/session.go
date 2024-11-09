package session

import (
	"errors"
	"time"

	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

const StatusClosed Status = "closed"

var ErrSessionClosed = errors.New("Session already closed")

type Status string

type Service struct {
	ServiceID      id.ID
	ProfessionalID id.ID
}

type Session struct {
	ID            id.ID
	AppointmentID id.ID
	Status        Status
	CloseTime     time.Time
	Services      []Service
}

func (s *Session) Close(services []Service) error {

	if s.IsClosed() {
		return ErrSessionClosed
	}

	s.CloseTime = time.Now()
	s.Status = StatusClosed

	for _, v := range services {
		s.Services = append(s.Services, v)
	}

	return nil
}

func (s *Session) IsClosed() bool {
	return s.Status == StatusClosed
}

func (s Session) GetID() id.ID {
	return s.ID
}
