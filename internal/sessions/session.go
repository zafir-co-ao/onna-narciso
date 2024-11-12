package sessions

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"

	"errors"
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

type Appointment struct {
	ID               nanoid.ID
	CustomerID       nanoid.ID
	CustomerName     string
	ProfessionalID   nanoid.ID
	ProfessionalName string
	ServiceID        nanoid.ID
	ServiceName      string
}

type Status string

type SessionService struct {
	ServiceID        nanoid.ID
	ServiceName      string
	ProfessionalID   nanoid.ID
	ProfessionalName string
}

type Session struct {
	ID            nanoid.ID
	AppointmentID nanoid.ID
	Status        Status
	Date          date.Date
	CheckinTime   hour.Hour
	StartTime     hour.Hour
	CloseTime     hour.Hour
	Services      []SessionService
	CustomerID    nanoid.ID
	CustomerName  string
}

func (s *Session) Start() error {

	if s.IsStarted() {
		return ErrSessionStarted
	}

	s.StartTime = hour.Now()
	s.Status = StatusStarted
	return nil
}

func (s *Session) Close(services []SessionService) error {

	if s.IsClosed() {
		return ErrSessionClosed
	}

	s.CloseTime = hour.Now()
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

func (s *Session) addServices(services []SessionService) {

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
