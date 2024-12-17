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
	ErrSessionStarted       = errors.New("session already started")
	ErrSessionClosed        = errors.New("session already closed")
	ErrInvalidCheckinDate   = errors.New("invalid checkin date")
	ErrInvoiceNotBeIssued   = errors.New("invoice not be issued")
	ErrInvalidStatusToStart = errors.New("invalid status to start")
	ErrInvalidStatusToClose = errors.New("invalid status to close")
)

type Appointment struct {
	ID               nanoid.ID
	CustomerID       nanoid.ID
	CustomerName     string
	ProfessionalID   nanoid.ID
	ProfessionalName string
	ServiceID        nanoid.ID
	ServiceName      string
	Date             date.Date
	Closed           bool
	Canceled         bool
}

func (a *Appointment) IsCanceled() bool {
	return a.Canceled
}

func (a *Appointment) IsClosed() bool {
	return a.Closed
}

func (a *Appointment) ValidCheckinDate() bool {
	return a.Date == date.Today()
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

	if !s.IsCheckedIn() {
		return ErrInvalidStatusToStart
	}

	s.StartTime = hour.Now()
	s.Status = StatusStarted
	return nil
}

func (s *Session) Close(services []SessionService) error {

	if s.IsClosed() {
		return ErrSessionClosed
	}

	if !s.IsStarted() {
		return ErrInvalidStatusToClose
	}

	s.CloseTime = hour.Now()
	s.Status = StatusClosed
	s.addServices(services)

	return nil
}

func (s *Session) IsCheckedIn() bool {
	return s.Status == StatusCheckedIn
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

	s.Services = append(s.Services, services...)
}

func (s Session) GetID() nanoid.ID {
	return s.ID
}
