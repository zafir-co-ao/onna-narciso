package scheduling

import (
	"github.com/zafir-co-ao/onna-narciso/internal/shared/event"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/id"
)

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

func (u *appointmentCancelerImpl) Execute(appointmentId string) error {
	a, err := u.repo.FindByID(id.NewID(appointmentId))
	if err != nil {
		return err
	}

	err = a.Cancel()
	if err != nil {
		return err
	}

	u.repo.Save(a)

	e := event.New(
		EventAppointmentCanceled,
		event.WithHeader(event.HeaderAggregateID, a.ID.String()),
	)

	u.bus.Publish(e)

	return nil
}
