package session

import (
	"github.com/kindalus/godx/pkg/nanoid"

	"errors"
	"time"
)

const StatusClosed Status = "closed"

var ErrSessionClosed = errors.New("Session already closed")

type Status string

type Service struct {
	ServiceID      nanoid.ID
	ProfessionalID nanoid.ID
}

type Session struct {
	ID            nanoid.ID
	AppointmentID nanoid.ID
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

func (s Session) GetID() nanoid.ID {
	return s.ID
}
