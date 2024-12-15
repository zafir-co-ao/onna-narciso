package sessions

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
)

const EventSessionCheckedIn = "EventSessionCheckedIn"

type Creator interface {
	Create(appointmentID string) (SessionOutput, error)
}

type creatorImpl struct {
	repo Repository
	bus  event.Bus
	aacl AppointmentsACL
}

func NewSessionCreator(repo Repository, bus event.Bus, aacl AppointmentsACL) Creator {
	return &creatorImpl{repo, bus, aacl}
}

func (u *creatorImpl) Create(appointmentID string) (SessionOutput, error) {
	a, err := u.aacl.FindByID(nanoid.ID(appointmentID))
	if err != nil {
		return SessionOutput{}, err
	}

	if !a.ValidCheckinDate() {
		return SessionOutput{}, ErrInvalidCheckinDate
	}

	if a.IsCanceled() {
		return SessionOutput{}, ErrAppointmentCanceled
	}

	if a.IsClosed() {
		return SessionOutput{}, ErrAppointmentClosed
	}

	s := NewSessionBuilder().
		WithAppointmentID(nanoid.ID(appointmentID)).
		WithCustomer(a.CustomerID, a.CustomerName).
		WithService(a.ServiceID, a.ServiceName, a.ProfessionalID, a.ProfessionalName).
		Build()

	err = u.repo.Save(s)
	if err != nil {
		return SessionOutput{}, err
	}

	p := struct{ AppointmentID string }{
		AppointmentID: s.AppointmentID.String(),
	}

	evt := event.New(EventSessionCheckedIn,
		event.WithHeader(event.HeaderAggregateType, "session.Session"),
		event.WithHeader(event.HeaderAggregateID, s.ID.String()),
		event.WithPayload(p),
	)

	u.bus.Publish(evt)

	return toSessionOutput(s), nil
}
