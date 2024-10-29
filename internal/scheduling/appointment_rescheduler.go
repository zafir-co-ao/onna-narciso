package scheduling

import "github.com/zafir-co-ao/onna-narciso/internal/shared/event"

const EventAppointmentRescheduled = "EventAppointmentRescheduled"

type AppointmentReschedulerInput struct {
	ID        string
	Date      string
	StartHour string
	Duration  int
}

type AppointmentRescheduler interface {
	Execute(i AppointmentReschedulerInput) (AppointmentOutput, error)
}

type appointmentRescheduler struct {
	repo AppointmentRepository
	bus  event.Bus
}

func NewAppointmentRescheduler(r AppointmentRepository, b event.Bus) AppointmentRescheduler {
	return &appointmentRescheduler{repo: r, bus: b}
}

func (r *appointmentRescheduler) Execute(i AppointmentReschedulerInput) (AppointmentOutput, error) {
	a, err := r.repo.FindByID(NewID(i.ID))
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	err = a.Reschedule(i.Date, i.StartHour, i.Duration)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	appointments, err := r.repo.FindByDateAndStatus(Date(i.Date), StatusScheduled)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	if !VerifyAvailability(a, appointments) {
		return EmptyAppointmentOutput, ErrBusyTime
	}

	r.repo.Save(a)

	e := event.New(
		EventAppointmentRescheduled,
		event.WithHeader(event.HeaderAggregateID, a.ID.Value()),
	)

	r.bus.Publish(e)

	return buildOutput(a), nil
}

type payload struct {
	Date      string
	StartHour string
}
