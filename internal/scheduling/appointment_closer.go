package scheduling

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
)

const EventAppointmentClosed = "EventAppointmentClosed"

type AppointmentCloser interface {
	Close(id string) error
}

type appointmentCloserImpl struct {
	repo AppointmentRepository
	bus  event.Bus
}

func NewAppointmentCloser(repo AppointmentRepository, bus event.Bus) AppointmentCloser {
	return &appointmentCloserImpl{repo, bus}
}

func (u *appointmentCloserImpl) Close(id string) error {
	a, err := u.repo.FindByID(nanoid.ID(id))
	if err != nil {
		return err
	}

	err = a.Close()
	if err != nil {
		return err
	}

	err = u.repo.Save(a)
	if err != nil {
		return err
	}

	e := event.New(
		EventAppointmentClosed,
		event.WithHeader(event.HeaderAggregateID, a.ID.String()),
		event.WithPayload(a),
	)

	u.bus.Publish(e)

	return nil
}
