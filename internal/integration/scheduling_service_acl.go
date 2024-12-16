package integration

import (
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

type SchedulingServiceACL interface {
	CloseAppointment(id nanoid.ID) error
	FindAppointmentByID(id nanoid.ID) (scheduling.AppointmentOutput, error)
}

type internalSchedulingServiceACL struct {
	closer scheduling.AppointmentCloser
	finder scheduling.AppointmentFinder
}

func NewInternalSchedulingServiceACL(finder scheduling.AppointmentFinder,
	closer scheduling.AppointmentCloser) SchedulingServiceACL {

	return &internalSchedulingServiceACL{
		closer: closer,
		finder: finder,
	}
}

func (s *internalSchedulingServiceACL) CloseAppointment(id nanoid.ID) error {
	return s.closer.Close(id.String())
}

func (s *internalSchedulingServiceACL) FindAppointmentByID(id nanoid.ID) (scheduling.AppointmentOutput, error) {
	return s.finder.FindByID(id.String())
}
