package sessions

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
	aacl AppointmentsACL
}

func NewSessionCreator(r Repository, b event.Bus, aacl AppointmentsACL) Creator {
	return &creatorImpl{repo: r, bus: b, aacl: aacl}
}

func (c *creatorImpl) Create(appointmentID string) (CreatorOutput, error) {
	a, err := c.aacl.FindByID(nanoid.ID(appointmentID))
	if err != nil {
		return CreatorOutput{}, err
	}

	if a.IsCanceled() {
		return CreatorOutput{}, ErrAppointmentCanceled
	}

	if !a.ValidCheckinDate() {
		return CreatorOutput{}, ErrInvalidCheckinDate
	}

	s := NewSessionBuilder().
		WithAppointmentID(nanoid.ID(appointmentID)).
		WithCustomer(a.CustomerID, a.CustomerName).
		WithService(a.ServiceID, a.ServiceName, a.ProfessionalID, a.ProfessionalName).
		Build()

	err = c.repo.Save(s)
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
