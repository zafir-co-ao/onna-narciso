package scheduling

import "github.com/zafir-co-ao/onna-narciso/internal/shared/event"

const EventAppointmentCanceled = "EventAppointmentCanceled"

type AppointmentCanceler interface {
	Execute(id string) error
}

type appointmentCancelerImpl struct {
	repo AppointmentRepository
	bus  event.Bus
}

func NewAppointmentCanceler(repo AppointmentRepository, bus event.Bus) AppointmentCanceler {
	return &appointmentCancelerImpl{repo, bus}
}

func (u *appointmentCancelerImpl) Execute(id string) error {
	a, err := u.repo.FindByID(NewID(id))
	if err != nil {
		return err
	}

	err = a.Cancel()
	if err != nil {
		return err
	}

	u.repo.Save(a)

	e := event.New(EventAppointmentCanceled)
	u.bus.Publish(e)

	return nil
}
