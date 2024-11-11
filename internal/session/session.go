package session

import (
	"github.com/kindalus/godx/pkg/nanoid"

	"errors"
	"time"
)

const (
	StatusCheckedIn Status = "CheckedIn"
	StatusStarted   Status = "started"
	StatusClosed    Status = "closed"
)

var (
	ErrSessionStarted = errors.New("Session already started")
	ErrSessionClosed  = errors.New("Session already closed")
)

type Status string

type Service struct {
	ServiceID      nanoid.ID
	ProfessionalID nanoid.ID
}

type Session struct {
	ID            nanoid.ID
	AppointmentID nanoid.ID
	Status        Status
	StartTime     time.Time
	CloseTime     time.Time
	Services      []Service
}

func (s *Session) Start() error {

	if s.IsStarted() {
		return ErrSessionStarted
	}

	s.StartTime = time.Now()
	s.Status = StatusStarted
	return nil
}

func (s *Session) Close(services []Service) error {

	if s.IsClosed() {
		return ErrSessionClosed
	}

	s.CloseTime = time.Now()
	s.Status = StatusClosed
	s.addServices(services)

	return nil
}

func (s *Session) IsStarted() bool {
	return s.Status == StatusStarted
}

func (s *Session) IsClosed() bool {
	return s.Status == StatusClosed
}

func (s *Session) addServices(services []Service) {

	if len(services) == 0 {
		return
	}

	for _, svc := range services {
		s.Services = append(s.Services, svc)
	}
}

func (s Session) GetID() nanoid.ID {
	return s.ID
}
