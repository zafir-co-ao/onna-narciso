package scheduling

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"
)

const EventAppointmentRescheduled = "EventAppointmentRescheduled"

type AppointmentReschedulerInput struct {
	ID             string
	ProfessionalID string
	ServiceID      string
	Date           string
	Hour           string
	Duration       int
}

type AppointmentRescheduler interface {
	Reschedule(i AppointmentReschedulerInput) error
}

type appointmentRescheduler struct {
	repo AppointmentRepository
	pacl ProfessionalsACL
	sacl ServicesServiceACL
	bus  event.Bus
}

func NewAppointmentRescheduler(
	repo AppointmentRepository,
	pacl ProfessionalsACL,
	sacl ServicesServiceACL,
	bus event.Bus,
) AppointmentRescheduler {
	return &appointmentRescheduler{repo, pacl, sacl, bus}
}

func (u *appointmentRescheduler) Reschedule(i AppointmentReschedulerInput) error {
	a, err := u.repo.FindByID(nanoid.ID(i.ID))
	if err != nil {
		return err
	}

	p, err := u.pacl.FindProfessionalByID(nanoid.ID(i.ProfessionalID))
	if err != nil {
		return err
	}

	s, err := u.sacl.FindServiceByID(nanoid.ID(i.ServiceID))
	if err != nil {
		return err
	}

	if !p.ContainsService(s.ID) {
		return ErrInvalidService
	}

	d, err := date.New(i.Date)
	if err != nil {
		return err
	}

	if d.Before() {
		return date.ErrDateInPast
	}

	h, err := hour.New(i.Hour)
	if err != nil {
		return err
	}

	a = NewAppointmentBuilder().
		WithAppointmentID(a.ID).
		WithCustomer(a.CustomerID, a.CustomerName).
		WithProfessional(p.ID, p.Name).
		WithService(s.ID, s.Name).
		WithDuration(duration.Duration(i.Duration)).
		WithStatus(a.Status).
		WithDate(d).
		WithHour(h).
		MustBuild()

	err = a.Reschedule()
	if err != nil {
		return err
	}

	appointments, err := u.repo.FindActivesByDateAndProfessional(d, p.ID)
	if err != nil {
		return err
	}

	if AppointmentsInterceptAny(a, appointments) {
		return ErrBusyTime
	}

	err = u.repo.Save(a)
	if err != nil {
		return err
	}

	e := event.New(
		EventAppointmentRescheduled,
		event.WithHeader(event.HeaderAggregateID, a.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return nil
}
