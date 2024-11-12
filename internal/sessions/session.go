package sessions

import (
	"github.com/kindalus/godx/pkg/nanoid"

	"errors"
	"time"
)

const (
	StatusCheckedIn Status = "CheckedIn"
	StatusStarted   Status = "Started"
	StatusClosed    Status = "Closed"
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
	CheckinTime   time.Time
	StartTime     time.Time
	CloseTime     time.Time
	Services      []Service
}

func NewSession(appointmentID nanoid.ID) Session {
	return Session{
		ID:            nanoid.New(),
		AppointmentID: appointmentID,
		Status:        StatusCheckedIn,
		CheckinTime:   time.Now(),
	}
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
