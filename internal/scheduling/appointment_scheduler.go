package scheduling

import (
	"github.com/kindalus/godx/pkg/event"
	"github.com/kindalus/godx/pkg/nanoid"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/date"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/duration"
	"github.com/zafir-co-ao/onna-narciso/internal/shared/hour"
)

const EventAppointmentScheduled = "EventAppointmentScheduled"

type AppointmentSchedulerInput struct {
	ID             string
	ProfessionalID string
	CustomerID     string
	CustomerName   string
	CustomerPhone  string
	ServiceID      string
	Date           string
	Hour           string
	Duration       int
}

type AppointmentScheduler interface {
	Schedule(i AppointmentSchedulerInput) (AppointmentOutput, error)
}

type appointmentScedulerImpl struct {
	repo AppointmentRepository
	sacl ServicesServiceACL
	cacl CRMServiceACL
	pacl HRServiceACL
	bus  event.Bus
}

func NewAppointmentScheduler(
	repo AppointmentRepository,
	cacl CRMServiceACL,
	pacl HRServiceACL,
	sacl ServicesServiceACL,
	bus event.Bus,
) AppointmentScheduler {
	return &appointmentScedulerImpl{
		repo: repo,
		cacl: cacl,
		pacl: pacl,
		sacl: sacl,
		bus:  bus,
	}
}

func (u *appointmentScedulerImpl) Schedule(i AppointmentSchedulerInput) (AppointmentOutput, error) {
	p, err := u.pacl.FindProfessionalByID(nanoid.ID(i.ProfessionalID))
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	s, err := u.sacl.FindServiceByID(nanoid.ID(i.ServiceID))
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	c, err := u.findOrRegistrationCustomer(i)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	d, err := date.New(i.Date)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	if d.Before() {
		return EmptyAppointmentOutput, date.ErrDateInPast
	}

	h, err := hour.New(i.Hour)
	if err != nil {
		return EmptyAppointmentOutput, err
	}

	a := NewAppointmentBuilder().
		WithProfessional(p.ID, p.Name).
		WithCustomer(c.ID, c.Name).
		WithService(s.ID, s.Name).
		WithDate(d).
		WithHour(h).
		WithDuration(duration.Duration(i.Duration)).
		MustBuild()

	appointments, err := u.repo.FindActivesByDateAndProfessional(d, p.ID)
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
		EventAppointmentScheduled,
		event.WithHeader(event.HeaderAggregateID, a.ID.String()),
		event.WithPayload(i),
	)

	u.bus.Publish(e)

	return toAppointmentOutput(a), nil
}

func (u *appointmentScedulerImpl) findOrRegistrationCustomer(i AppointmentSchedulerInput) (Customer, error) {
	isEmptyCustomer := len(i.CustomerName) == 0 && len(i.CustomerPhone) == 0

	if len(i.CustomerID) == 0 && isEmptyCustomer {
		return Customer{}, ErrCustomerNotFound
	}

	if len(i.CustomerID) > 0 && isEmptyCustomer {
		return u.cacl.FindCustomerByID(nanoid.ID(i.CustomerID))
	}

	return u.cacl.RequestCustomerRegistration(i.CustomerName, i.CustomerPhone)
}
