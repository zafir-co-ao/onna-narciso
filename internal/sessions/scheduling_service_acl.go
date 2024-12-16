package sessions

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

type SchedulingServiceACL interface {
	FindAppointmentByID(i nanoid.ID) (Appointment, error)
}

type SchedulingServiceACLFunc func(i nanoid.ID) (Appointment, error)

func (f SchedulingServiceACLFunc) FindAppointmentByID(i nanoid.ID) (Appointment, error) {
	return f(i)
}

type internalSchedulingServiceACL struct {
	finder scheduling.AppointmentFinder
}

func NewInternalSchedulingServiceACL(finder scheduling.AppointmentFinder) SchedulingServiceACL {
	return &internalSchedulingServiceACL{finder}
}

func (s *internalSchedulingServiceACL) FindAppointmentByID(i nanoid.ID) (Appointment, error) {
	a, err := s.finder.FindByID(i.String())
	if err != nil {
		return Appointment{}, err
	}

	return Appointment{
		ID:        nanoid.ID(a.ID),
		ServiceID: nanoid.ID(a.ServiceID),
	}, nil
}
