package scheduling

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
)

const EventAppointmentRescheduled = "EventAppointmentRescheduled"

type AppointmentReschedulerInput struct {
	ID             string
	ProfessionalID string
	Date           string
	Hour           string
	Duration       int
}

type AppointmentRescheduler interface {
	Reschedule(i AppointmentReschedulerInput) (AppointmentOutput, error)
}

type appointmentRescheduler struct {
	repo AppointmentRepository
	bus  event.Bus
}

func NewAppointmentRescheduler(r AppointmentRepository, b event.Bus) AppointmentRescheduler {
	return &appointmentRescheduler{repo: r, bus: b}
}

func (u *appointmentRescheduler) Reschedule(i AppointmentReschedulerInput) (AppointmentOutput, error) {
	a, err := u.repo.FindByID(nanoid.ID(i.ID))
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	err = a.Reschedule(i.ProfessionalID, i.Date, i.Hour, i.Duration)
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
