package scheduling

import (
	"slices"

	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
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
	Reschedule(i AppointmentReschedulerInput) (AppointmentOutput, error)
}

type appointmentRescheduler struct {
	repo  AppointmentRepository
	pacl  ProfessionalsACL
	sacl  ServicesACL
	bus   event.Bus
	clock Clock
}

func NewAppointmentRescheduler(
	repo AppointmentRepository,
	pacl ProfessionalsACL,
	sacl ServicesACL,
	bus event.Bus,
	clock Clock,
) AppointmentRescheduler {
	return &appointmentRescheduler{repo, pacl, sacl, bus, clock}
}

func (u *appointmentRescheduler) Reschedule(i AppointmentReschedulerInput) (AppointmentOutput, error) {
	a, err := u.repo.FindByID(nanoid.ID(i.ID))
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	p, err := u.pacl.FindProfessionalByID(nanoid.ID(i.ProfessionalID))
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	s, err := u.sacl.FindServiceByID(nanoid.ID(i.ServiceID))
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	if !slices.Contains(p.ServicesIDS, s.ID) {
		return EmptyAppointmentOutput, ErrInvalidService
	}

	d, err := date.New(i.Date)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	if d.Before(u.clock.Today()) {
		return EmptyAppointmentOutput, ErrScheduleInPast
	}

	h, err := hour.New(i.Hour)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	a = NewAppointmentBuilder().
		WithAppointmentID(a.ID).
		WithCustomer(a.CustomerID, a.CustomerName).
		WithProfessional(p.ID, p.Name).
		WithService(s.ID, s.Name).
		WithDuration(i.Duration).
		WithStatus(a.Status).
		WithDate(d).
		WithHour(h).
		Build()

	err = a.Reschedule()
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	appointments, err := u.repo.FindByDateStatusAndProfessional(date.Date(i.Date), StatusScheduled, a.ProfessionalID)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	if AppointmentsInterceptAny(a, appointments) {
		return EmptyAppointmentOutput, ErrBusyTime
	}

	err = u.repo.Save(a)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	e := event.New(
		EventAppointmentRescheduled,
		event.WithHeader(event.HeaderAggregateID, a.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return toAppointmentOutput(a), nil
}
