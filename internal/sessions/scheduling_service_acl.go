package sessions

import (
	"errors"

	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

var (
	ErrAppointmentNotFound = errors.New("appointment not found")
	ErrAppointmentCanceled = errors.New("appointment canceled")
	ErrAppointmentClosed   = errors.New("appointment closed")
)

type SchedulingServiceACL interface {
	FindAppointmentByID(id nanoid.ID) (Appointment, error)
}

type SchedulingServiceACLFunc func(id nanoid.ID) (Appointment, error)

func (f SchedulingServiceACLFunc) FindAppointmentByID(id nanoid.ID) (Appointment, error) {
	return f(id)
}

type internalSchedulingServiceACL struct {
	finder scheduling.AppointmentFinder
}

func NewInternalSchedulingServiceACL(finder scheduling.AppointmentFinder) SchedulingServiceACL {
	return &internalSchedulingServiceACL{finder}
}

func (s *internalSchedulingServiceACL) FindAppointmentByID(id nanoid.ID) (Appointment, error) {
	a, err := s.finder.FindByID(id.String())
	if err != nil {
		return Appointment{}, err
	}

	return Appointment{
		ID:               nanoid.ID(a.ID),
		ServiceID:        nanoid.ID(a.ServiceID),
		ServiceName:      a.ServiceName,
		CustomerID:       nanoid.ID(a.CustomerID),
		CustomerName:     a.CustomerName,
		ProfessionalID:   nanoid.ID(a.ProfessionalID),
		ProfessionalName: a.ProfessionalName,
		Date:             date.Date(a.Date),
		Closed:           a.Status == string(scheduling.StatusClosed),
		Canceled:         a.Status == string(scheduling.StatusCanceled),
	}, nil
}
