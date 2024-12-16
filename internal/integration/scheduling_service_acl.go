package integration

import (
	"github.com/zafir-co-ao/onna-narciso/internal/scheduling"
)

type SchedulingServiceACL interface {
	CloseAppointment(id string) error
	GetAppointment(id string) (scheduling.AppointmentOutput, error)
}

type internalSchedulingServiceACL struct {
	closer scheduling.AppointmentCloser
	getter scheduling.AppointmentGetter
}

func NewInternalSchedulingServiceACL(getter scheduling.AppointmentGetter, closer scheduling.AppointmentCloser) SchedulingServiceACL {
	return &internalSchedulingServiceACL{closer: closer}
}

func (s *internalSchedulingServiceACL) CloseAppointment(id string) error {
	return s.closer.Close(id)
}

func (s *internalSchedulingServiceACL) GetAppointment(id string) (scheduling.AppointmentOutput, error) {
	return s.getter.Get(id)
}
