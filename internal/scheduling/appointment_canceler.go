package scheduling

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
)

const EventAppointmentCanceled = "EventAppointmentCanceled"

type AppointmentCanceler interface {
	Cancel(id string) error
}

type appointmentCancelerImpl struct {
	repo AppointmentRepository
	bus  event.Bus
}

func NewAppointmentCanceler(repo AppointmentRepository, bus event.Bus) AppointmentCanceler {
	return &appointmentCancelerImpl{repo, bus}
}

func (u *appointmentCancelerImpl) Cancel(appointmentId string) error {
	a, err := u.repo.FindByID(nanoid.ID(appointmentId))
	if err != nil {
		return err
	}

	err = a.Cancel()
	if err != nil {
		return err
	}

	err = u.repo.Save(a)
	if err != nil {
		return err
	}

	e := event.New(
		EventAppointmentCanceled,
		event.WithHeader(event.HeaderAggregateID, a.ID.String()),
	)

	u.bus.Publish(e)

	return nil
}
