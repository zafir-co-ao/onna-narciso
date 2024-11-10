package session

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
)

type CreatorOutput struct {
	ID            string
	AppointmentID string
}

type Creator interface {
	Create(appointmentID string) (CreatorOutput, error)
}

type creatorImpl struct {
	repo Repository
	bus  event.Bus
}

func NewSessionCreator(r Repository, b event.Bus) Creator {
	return &creatorImpl{repo: r, bus: b}
}

func (c *creatorImpl) Create(appointmentID string) (CreatorOutput, error) {

	_id := nanoid.New()

	s := Session{
		ID:            _id,
		AppointmentID: nanoid.ID(appointmentID),
	}

	err := c.repo.Save(s)
	if err != nil {
		return CreatorOutput{}, err
	}

	c.bus.Publish(event.New("SessionCheckedIn",
		event.WithHeader(event.HeaderAggregateID, s.ID.String()),
		event.WithPayload(struct{ AppointmentID string }{
			AppointmentID: s.AppointmentID.String(),
		}),
	))

	return CreatorOutput{
		ID:            s.ID.String(),
		AppointmentID: s.AppointmentID.String(),
	}, nil
}
